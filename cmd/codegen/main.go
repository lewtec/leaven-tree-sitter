package main

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var (
	// treeSitterPath is set in run() via TREE_SITTER_PATH or workspaced cache.
	treeSitterPath   string
	defaultOutputDir = "."
	keepTemp         bool
	clangBin         string
	onlyLangs        []string
)

var rootCmd = &cobra.Command{
	Use:   "codegen",
	Short: "Transpile tree-sitter C to Go via clang IR + leaven",
	Long: `Transpile tree-sitter core and grammars with clang -emit-llvm and leaven.

No ccgo. Prefer clang 22; pin with mise conda:clang@22.
Use --only to limit grammar units (e.g. --only=json,python). With no --only,
all discovered units under third-party/tree-sitter-*/ are attempted.

Writes grammar/core.go and grammar/<lang>/grammar.go (raw leaven output).`,
	RunE: run,
}

func env(key, defaultValue string) string {
	if ret := os.Getenv(key); ret != "" {
		return ret
	}
	return defaultValue
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&keepTemp, "keep-temp", "k", false, "Keep temporary files for debugging")
	rootCmd.Flags().StringVar(&clangBin, "clang", env("CC", "clang"), "clang binary (use 22.x for leaven)")
	rootCmd.Flags().StringSliceVar(&onlyLangs, "only", nil, "Limit to these language ids (repeatable or comma-separated)")

	rootCmd.AddCommand(&cobra.Command{
		Use:   "print-tree-sitter-path",
		Short: "Print workspaced cache path for core tree-sitter (fetch if missing)",
		RunE: func(cmd *cobra.Command, args []string) error {
			p, err := resolveTreeSitterPath()
			if err != nil {
				return err
			}
			fmt.Println(p)
			return nil
		},
	})

	modulesCmd := &cobra.Command{
		Use:   "modules",
		Short: "Write nested go.mod files and go.work for grammar packages",
		Long: `Write grammar/go.mod, grammar/<lang>/go.mod (with local replace
directives), and go.work. Optionally run go work sync / go mod tidy.

Does not transpile C sources.`,
		RunE: runModules,
	}
	modulesCmd.Flags().Bool("tidy", false, "Run go work sync and go mod tidy for each module")
	rootCmd.AddCommand(modulesCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func runModules(cmd *cobra.Command, args []string) error {
	tidy, err := cmd.Flags().GetBool("tidy")
	if err != nil {
		return fmt.Errorf("failed to get --tidy flag: %w", err)
	}
	slog.Info("writing grammar modules", "dir", defaultOutputDir, "tidy", tidy)
	if err := ensureGrammarModules(defaultOutputDir); err != nil {
		return err
	}
	if tidy {
		slog.Info("tidying workspace modules")
		if err := tidyGrammarModules(defaultOutputDir); err != nil {
			return err
		}
	}
	return nil
}

func run(cmd *cobra.Command, args []string) error {
	var err error
	treeSitterPath, err = resolveTreeSitterPath()
	if err != nil {
		return fmt.Errorf("tree-sitter source: %w", err)
	}

	lt := &LeavenTranspiler{
		TreeSitterPath: treeSitterPath,
		KeepTemp:       keepTemp,
		Clang:          clangBin,
	}

	grammarDir := filepath.Join(defaultOutputDir, "grammar")
	slog.Info("transpiling tree-sitter core via leaven", "path", treeSitterPath, "clang", clangBin)
	if err := lt.TranspileCore(cmd.Context(), grammarDir); err != nil {
		// Core IR often uses atomicrmw / intrinsics leaven cannot parse yet.
		// Still migrate every grammar; core remains a known follow-up.
		slog.Error("leaven core failed; continuing with grammars only", "error", err)
	}

	units, err := discoverGrammarUnits("third-party/tree-sitter-*")
	if err != nil {
		return err
	}
	if len(units) == 0 {
		return fmt.Errorf("no grammar units under third-party/tree-sitter-*/src/parser.c; run `mise run grammars:sync` first")
	}
	if len(onlyLangs) > 0 {
		want := map[string]bool{}
		for _, a := range onlyLangs {
			want[normalizeGrammarName(a)] = true
		}
		filtered := units[:0]
		for _, u := range units {
			if want[u.Name] {
				filtered = append(filtered, u)
			}
		}
		units = filtered
		if len(units) == 0 {
			return fmt.Errorf("no units matched --only=%v", onlyLangs)
		}
	}
	slog.Info("discovered grammar units", "count", len(units))

	var summaryWriter io.Writer
	if summaryPath := os.Getenv("GITHUB_STEP_SUMMARY"); summaryPath != "" {
		summaryFile, err := os.OpenFile(summaryPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			slog.Warn("failed to open GITHUB_STEP_SUMMARY", "path", summaryPath, "error", err)
		} else {
			summaryWriter = summaryFile
			defer summaryFile.Close()
		}
	} else {
		var buf bytes.Buffer
		defer fmt.Println(buf.String())
		summaryWriter = &buf
	}

	fmt.Fprintf(summaryWriter, "## Grammar Codegen Summary (leaven)\n\n")
	var failed int
	for i, unit := range units {
		slog.Info("transpiling grammar", "index", i+1, "total", len(units), "grammar", unit.Name, "path", unit.Path)
		if err := lt.TranspileGrammar(cmd.Context(), unit.Path, unit.Name, grammarDir); err != nil {
			slog.Warn("failed to transpile grammar, skipping", "grammar", unit.Name, "path", unit.Path, "error", err)
			fmt.Fprintf(summaryWriter, "- `%s`: ❌ (%v)\n", unit.Name, err)
			failed++
			continue
		}
		fmt.Fprintf(summaryWriter, "- `%s`: ✅\n", unit.Name)
	}

	slog.Info("updating languages registry in cmd/parse/languages.go")
	if err := updateLanguagesGo(defaultOutputDir); err != nil {
		return fmt.Errorf("failed to update languages registry: %w", err)
	}

	slog.Info("writing per-grammar go.mod and go.work")
	if err := ensureGrammarModules(defaultOutputDir); err != nil {
		return fmt.Errorf("failed to write grammar modules: %w", err)
	}

	if failed > 0 {
		slog.Warn("skipped failed grammars", "failed", failed, "total", len(units))
		fmt.Fprintf(summaryWriter, "\n%d/%d grammars failed (skipped).\n", failed, len(units))
	}
	if allGrammarsFailed(failed, len(units)) {
		return fmt.Errorf("leaven: %d/%d grammars failed", failed, len(units))
	}
	return nil
}

func allGrammarsFailed(failed, total int) bool {
	return total > 0 && failed == total
}

func updateLanguagesGo(outputDir string) error {
	grammarDir := filepath.Join(outputDir, "grammar")
	languages, err := listGrammarLangs(grammarDir)
	if err != nil {
		return err
	}

	moduleName := "github.com/lewtec/leaven-tree-sitter"

	var sb strings.Builder
	sb.WriteString("// Code generated by codegen (updateLanguagesGo); DO NOT EDIT.\n\n")
	sb.WriteString("package main\n")
	if len(languages) > 0 {
		sb.WriteString("\nimport (\n")
		for _, lang := range languages {
			fmt.Fprintf(&sb, "\t_ \"%s/grammar/%s\"\n", moduleName, lang)
		}
		sb.WriteString(")\n")
	} else {
		sb.WriteString("\n// No grammar packages on disk yet. Cook PRs fill this file.\n")
	}

	targetFile := filepath.Join(outputDir, "cmd", "parse", "languages.go")
	return os.WriteFile(targetFile, []byte(sb.String()), 0644)
}
