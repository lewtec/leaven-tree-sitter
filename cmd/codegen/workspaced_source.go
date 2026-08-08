package main

import (
	"archive/tar"
	"compress/gzip"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Workspaced github source cache layout (matches lucasew/workspaced sourcecache):
//
//	~/.cache/workspaced/sources/github/sha256("v4:repo:OWNER/REPO@REF")
//
// Cache key uses the cue input version (default HEAD), not the lock digest.
// The lock digest pins which tarball is fetched into that directory.

// workspacedGithubSource is one workspaced github input: cue `from` + `version`,
// optional env override, and a marker file that means the tree is ready.
type workspacedGithubSource struct {
	Repo        string // owner/name (e.g. tree-sitter/tree-sitter)
	Version     string // cue version; empty means HEAD
	Marker      string // relative path that must exist when ready
	EnvOverride string // if set and non-empty, use that path instead of the cache
}

// treeSitterSource is the core C library. Keep Version in sync with
// #tree_sitter.version in workspaced.cue.
var treeSitterSource = workspacedGithubSource{
	Repo:        "tree-sitter/tree-sitter",
	Version:     "HEAD",
	Marker:      "lib/src/lib.c",
	EnvOverride: "TREE_SITTER_PATH",
}

func (s workspacedGithubSource) version() string {
	v := strings.TrimSpace(s.Version)
	if v == "" {
		return "HEAD"
	}
	return v
}

func (s workspacedGithubSource) repo() string {
	return strings.Trim(strings.TrimSpace(s.Repo), "/")
}

// cacheKey matches workspaced github.Source.CacheKey for repo+ref (no URL form).
func (s workspacedGithubSource) cacheKey() string {
	return "v4:repo:" + s.repo() + "@" + s.version()
}

// CachePath is the on-disk workspaced cache directory for this input.
func (s workspacedGithubSource) CachePath() (string, error) {
	sum := sha256.Sum256([]byte(s.cacheKey()))
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("user home: %w", err)
	}
	return filepath.Join(home, ".cache", "workspaced", "sources", "github", hex.EncodeToString(sum[:])), nil
}

// Resolve returns EnvOverride (if set) or the workspaced cache path, ensuring
// the locked tree is present on disk.
func (s workspacedGithubSource) Resolve() (string, error) {
	if s.EnvOverride != "" {
		if p := strings.TrimSpace(os.Getenv(s.EnvOverride)); p != "" {
			if err := s.checkReady(p); err != nil {
				return "", err
			}
			return p, nil
		}
	}
	return s.ensure()
}

func (s workspacedGithubSource) checkReady(root string) error {
	if s.Marker == "" {
		if st, err := os.Stat(root); err != nil || !st.IsDir() {
			return fmt.Errorf("%s root %s: not a directory", s.repo(), root)
		}
		return nil
	}
	mark := filepath.Join(root, s.Marker)
	if _, err := os.Stat(mark); err != nil {
		return fmt.Errorf("%s root %s: %w (need %s)", s.repo(), root, err, s.Marker)
	}
	return nil
}

// ensure returns the cache path, fetching the locked commit when missing.
func (s workspacedGithubSource) ensure() (string, error) {
	cache, err := s.CachePath()
	if err != nil {
		return "", err
	}
	if err := s.checkReady(cache); err == nil {
		return cache, nil
	}

	digest, err := lockDigestForGithubSource(s.repo())
	if err != nil {
		return "", fmt.Errorf("%s cache miss at %s: %w\nrun: mise run grammars:lock", s.repo(), cache, err)
	}
	if err := fetchGithubTarballToCache(codeloadURL(s.repo(), digest), cache); err != nil {
		return "", fmt.Errorf("fetch %s into workspaced cache: %w", s.repo(), err)
	}
	if err := s.checkReady(cache); err != nil {
		return "", err
	}
	return cache, nil
}

func codeloadURL(repo, digest string) string {
	return fmt.Sprintf("https://codeload.github.com/%s/tar.gz/%s", repo, digest)
}

// resolveTreeSitterPath is the codegen entrypoint for core tree-sitter sources.
func resolveTreeSitterPath() (string, error) {
	return treeSitterSource.Resolve()
}

// lockDigestForGithubSource reads currentDigest for github:OWNER/REPO from
// workspaced.lock.json in the current working directory (or parents).
func lockDigestForGithubSource(repo string) (string, error) {
	lockPath, err := findUpFile("workspaced.lock.json")
	if err != nil {
		return "", err
	}
	b, err := os.ReadFile(lockPath)
	if err != nil {
		return "", err
	}
	var lock struct {
		Dependencies []struct {
			Kind          string `json:"kind"`
			Ref           string `json:"ref"`
			DepName       string `json:"depName"`
			CurrentDigest string `json:"currentDigest"`
		} `json:"dependencies"`
	}
	if err := json.Unmarshal(b, &lock); err != nil {
		return "", fmt.Errorf("parse %s: %w", lockPath, err)
	}
	wantRef := "github:" + repo
	for _, d := range lock.Dependencies {
		if d.Kind != "source" {
			continue
		}
		if d.Ref == wantRef || d.DepName == repo {
			if dig := strings.TrimSpace(d.CurrentDigest); dig != "" {
				return dig, nil
			}
		}
	}
	return "", fmt.Errorf("no locked source %s in %s (mise run grammars:lock)", wantRef, lockPath)
}

// findUpFile walks from the working directory upward looking for name.
func findUpFile(name string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	dir := wd
	for {
		p := filepath.Join(dir, name)
		if st, err := os.Stat(p); err == nil && !st.IsDir() {
			return p, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return "", fmt.Errorf("%s not found from %s", name, wd)
}

// fetchGithubTarballToCache downloads a GitHub codeload tarball into the
// workspaced cache directory layout (strip top-level repo-sha/ prefix).
func fetchGithubTarballToCache(url, dest string) error {
	if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
		return err
	}
	tmp := dest + ".tmp"
	_ = os.RemoveAll(tmp)
	if err := os.MkdirAll(tmp, 0o755); err != nil {
		return err
	}
	defer func() { _ = os.RemoveAll(tmp) }()

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", "ccgo-tree-sitter-codegen (+https://github.com/lewtec/leaven-tree-sitter)")
	if tok := strings.TrimSpace(os.Getenv("GITHUB_TOKEN")); tok != "" {
		req.Header.Set("Authorization", "Bearer "+tok)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("GET %s: %s", url, resp.Status)
	}

	h := sha256.New()
	if err := extractGithubTarGz(io.TeeReader(resp.Body, h), tmp); err != nil {
		return err
	}
	hash := hex.EncodeToString(h.Sum(nil))
	meta, err := json.Marshal(map[string]string{"url": url, "hash": hash})
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(tmp, ".workspaced-source-meta.json"), meta, 0o644); err != nil {
		return err
	}
	return installDir(tmp, dest)
}

// installDir moves tmp into dest (rename, or copy on cross-device failure).
// On success tmp is gone; caller must not rely on tmp existing afterward.
func installDir(tmp, dest string) error {
	_ = os.RemoveAll(dest)
	if err := os.Rename(tmp, dest); err == nil {
		return nil
	} else if err2 := copyDir(tmp, dest); err2 != nil {
		return fmt.Errorf("install cache dir: rename: %v; copy: %w", err, err2)
	}
	_ = os.RemoveAll(tmp)
	return nil
}

func extractGithubTarGz(r io.Reader, destDir string) error {
	gzr, err := gzip.NewReader(r)
	if err != nil {
		return err
	}
	defer gzr.Close()
	tr := tar.NewReader(gzr)
	destClean := filepath.Clean(destDir)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		// Strip GitHub's top-level "repo-sha/" component.
		name := strings.TrimPrefix(hdr.Name, "./")
		slash := strings.IndexByte(name, '/')
		if slash < 0 {
			continue
		}
		rel := name[slash+1:]
		if rel == "" {
			continue
		}
		target := filepath.Join(destDir, filepath.FromSlash(rel))
		if !pathUnder(destClean, target) {
			return fmt.Errorf("illegal path in tarball: %s", hdr.Name)
		}
		switch hdr.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(target, 0o755); err != nil {
				return err
			}
		case tar.TypeReg:
			if err := writeTarFile(target, os.FileMode(hdr.Mode)&0o755|0o644, tr); err != nil {
				return err
			}
		case tar.TypeSymlink:
			if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
				return err
			}
			_ = os.Remove(target)
			if err := os.Symlink(hdr.Linkname, target); err != nil {
				return err
			}
		}
	}
}

func pathUnder(root, target string) bool {
	root = filepath.Clean(root)
	target = filepath.Clean(target)
	if target == root {
		return true
	}
	sep := string(os.PathSeparator)
	return strings.HasPrefix(target+sep, root+sep)
}

func writeTarFile(target string, mode os.FileMode, r io.Reader) error {
	if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
		return err
	}
	f, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, mode)
	if err != nil {
		return err
	}
	_, copyErr := io.Copy(f, r)
	closeErr := f.Close()
	if copyErr != nil {
		return copyErr
	}
	return closeErr
}

func copyDir(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		out := filepath.Join(dst, rel)
		if info.IsDir() {
			return os.MkdirAll(out, 0o755)
		}
		if info.Mode()&os.ModeSymlink != 0 {
			link, err := os.Readlink(path)
			if err != nil {
				return err
			}
			_ = os.Remove(out)
			return os.Symlink(link, out)
		}
		in, err := os.Open(path)
		if err != nil {
			return err
		}
		defer in.Close()
		return writeTarFile(out, info.Mode(), in)
	})
}
