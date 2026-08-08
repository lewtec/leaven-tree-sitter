package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/lewtec/leaven-tree-sitter/grammar"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:           "parse <language> <file>",
	Short:         "Parse source files and print tree-sitter nodes",
	Args:          cobra.ExactArgs(2),
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(_ *cobra.Command, args []string) error {
		return run(args[0], args[1])
	},
}

var querySource string
var queryFile string

func init() {
	Command.Flags().StringVar(&querySource, "query", "", "Tree-sitter query source")
	Command.Flags().StringVar(&queryFile, "query-file", "", "Path to a file containing a tree-sitter query")
}

func main() {
	Command.SetUsageTemplate(Command.UsageTemplate() + "\nSupported languages:\n" + grammar.SupportedLanguages() + "\n")
	if err := Command.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(languageName, filename string) error {
	if querySource != "" && queryFile != "" {
		return fmt.Errorf("use either --query or --query-file, not both")
	}

	// Read file
	source, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	// Resolve language. Get can return (nil, true) if a nil Language was
	// registered; treat that like an unsupported name so we never SetLanguage(nil).
	lang, ok := grammar.Get(strings.ToLower(languageName))
	if !ok || lang == nil {
		return fmt.Errorf(
			"unsupported language: %s\nsupported languages: %s",
			languageName,
			grammar.SupportedLanguages(),
		)
	}

	// Create parser (GC-managed; no Delete required)
	parser := grammar.NewParser()
	if !parser.SetLanguage(lang) {
		return fmt.Errorf("failed to set language")
	}

	// Parse
	tree := parser.ParseBytes(source)

	query := querySource
	if queryFile != "" {
		queryBytes, err := os.ReadFile(queryFile)
		if err != nil {
			return fmt.Errorf("error reading query file: %w", err)
		}
		query = string(queryBytes)
	}

	// Convert parse tree to JSON
	root := tree.RootNode()
	if root.IsNull() {
		return fmt.Errorf("parse failed: tree root is null")
	}
	if query != "" {
		compiledQuery, err := grammar.NewQuery(lang, query)
		if err != nil {
			return err
		}
		matches := compiledQuery.ExecuteMatches(root, source)
		enc := json.NewEncoder(os.Stdout)
		for _, match := range matches {
			if err := enc.Encode(match); err != nil {
				return fmt.Errorf("failed to encode match output: %w", err)
			}
		}
		return nil
	}

	output := grammar.ParseOutput{
		Language: strings.ToLower(languageName),
		File:     filename,
		Root:     grammar.BuildParseNode(root, source, ""),
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(output); err != nil {
		return fmt.Errorf("failed to encode output: %w", err)
	}
	return nil
}
