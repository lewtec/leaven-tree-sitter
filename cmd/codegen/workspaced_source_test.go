package main

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWorkspacedGithubCachePath(t *testing.T) {
	s := workspacedGithubSource{Repo: "tree-sitter/tree-sitter", Version: "HEAD"}
	path, err := s.CachePath()
	if err != nil {
		t.Fatal(err)
	}
	sum := sha256.Sum256([]byte(s.cacheKey()))
	wantHash := hex.EncodeToString(sum[:])
	if !strings.HasSuffix(path, filepath.Join("sources", "github", wantHash)) {
		t.Fatalf("path = %s, want …/sources/github/%s", path, wantHash)
	}
	if !filepath.IsAbs(path) {
		t.Fatalf("path not absolute: %s", path)
	}
	if s.cacheKey() != "v4:repo:tree-sitter/tree-sitter@HEAD" {
		t.Fatalf("cacheKey = %q", s.cacheKey())
	}
}

func TestTreeSitterSourceMatchesCue(t *testing.T) {
	if treeSitterSource.Repo != "tree-sitter/tree-sitter" {
		t.Fatalf("repo = %q", treeSitterSource.Repo)
	}
	if treeSitterSource.Version != "HEAD" {
		t.Fatalf("version = %q (keep in sync with #tree_sitter.version)", treeSitterSource.Version)
	}
	if treeSitterSource.EnvOverride != "TREE_SITTER_PATH" {
		t.Fatalf("env = %q", treeSitterSource.EnvOverride)
	}
}

func TestLockDigestForGithubSource(t *testing.T) {
	root := findRepoRoot(t)
	if _, err := os.Stat(filepath.Join(root, "workspaced.lock.json")); err != nil {
		t.Skip("no workspaced.lock.json")
	}
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(root); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chdir(wd) })

	dig, err := lockDigestForGithubSource(treeSitterSource.repo())
	if err != nil {
		if strings.Contains(err.Error(), "no locked source") {
			t.Skip(err.Error())
		}
		t.Fatal(err)
	}
	if len(dig) < 7 {
		t.Fatalf("digest too short: %q", dig)
	}
}

func TestPathUnder(t *testing.T) {
	root := filepath.Join(string(filepath.Separator), "cache", "src")
	if !pathUnder(root, root) {
		t.Fatal("root should contain itself")
	}
	if !pathUnder(root, filepath.Join(root, "lib", "src")) {
		t.Fatal("child should be under root")
	}
	if pathUnder(root, filepath.Join(string(filepath.Separator), "cache", "other")) {
		t.Fatal("sibling must not be under root")
	}
	if pathUnder(root, filepath.Join(string(filepath.Separator), "cache", "src-evil")) {
		t.Fatal("prefix sibling must not be under root")
	}
}

func TestFindUpFile(t *testing.T) {
	root := findRepoRoot(t)
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	// Start from cmd/codegen so we walk up.
	if err := os.Chdir(filepath.Join(root, "cmd", "codegen")); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chdir(wd) })

	got, err := findUpFile("workspaced.cue")
	if err != nil {
		t.Fatal(err)
	}
	want := filepath.Join(root, "workspaced.cue")
	if filepath.Clean(got) != filepath.Clean(want) {
		t.Fatalf("got %s want %s", got, want)
	}
}
