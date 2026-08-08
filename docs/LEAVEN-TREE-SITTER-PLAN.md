# Plan: `lewtec/leaven-tree-sitter`

Clean port of the leaven tree-sitter Go stack from
`ccgo-tree-sitter` branch `20260806-try-codegen`. New repo. No ccgo.
No `modernc.org/libc`. No history transplant.

Reader: whoever implements Phase 1. After this doc you can land the
harness on `main` and dispatch a cook of the live `#grammar` set.

Source of truth: grilling 2026-08-07/08, closed Q12.

---

## Product shape

| Item | Decision |
|------|----------|
| Repo | https://github.com/lewtec/leaven-tree-sitter |
| Module path | `github.com/lewtec/leaven-tree-sitter` |
| Backend | clang 14 → LLVM IR → `go tool leaven` (lewtec/leaven pin) |
| Layout | One module for core (`…/grammar`) + one module per language (`…/grammar/<lang>`) |
| Scheme | Codegen → nested modules → `Register` → parse/fixtures API |

---

## What we commit

| Commit | Do not commit |
|--------|----------------|
| Hand API (`api.go`, `query.go`, `parse.go`, `registry.go`, `status.go`, …) | Original C under `third-party/tree-sitter-*` |
| `cmd/codegen` (leaven only) | `grammar/<lang>/` from a laptop cook |
| `grammar/core.go` (from local cook) | Generated lang `.go` until a cook PR merges |
| `workspaced.cue` + `workspaced.lock.json` | Full 190-lang lock |
| testdata + goldens (full origin set; unused until a lang is registered) | |
| mise / CI / `go.mod` / `go.work` | |

C sources: workspaced JIT fetch when you cook. The lockfile is the pin.

---

## Phase 1 scope

Focus is the **factory** and **core on `main`**. Languages are inventory:
uncomment → lock → cook PR. The product is the full set, not one language.

Day-one live `#grammar`: **php** (plus tree-sitter core). That is the first
smoke, not the destination. Every other language is a `//` comment.

Day-one lockfile pins tree-sitter core + php. Nothing else.

Codegen with no `--only` cooks **all placed / live** grammars. `--only`
is a local debug filter.

The old “16 corpus langs on day one” list is later expansion. Nested Go
modules per grammar.

---

## Local loop

Work on the harness and `core.go` on `main`.

1. Run codegen locally (clang 14 + leaven + JIT C). Use `--only=<lang>`
   to keep the loop short.
2. Require green local tests for core + the langs you cooked.
3. Keep or update `grammar/core.go`.
4. Discard every `grammar/<lang>/`. Do not push lang output from the laptop.
5. Push the harness + `core.go` when that loop is green.

Do not debug the first clang/PATH/leaven failure only on GitHub Actions.

---

## Codegen workflow

Trigger: **`workflow_dispatch`**, or **push to `main`** that changes
`workspaced.lock.json`. No cron. Concurrency group `codegen` with
`cancel-in-progress: false` — one cook at a time; later runs wait.

Runs once on ubuntu (clang 14 + leaven). One pure-Go output. No
GOOS/GOARCH codegen matrix.

Steps:

1. You decide local state is good.
2. Dispatch codegen, or push a lockfile change to `main`.
3. Job JIT-places C with `GITHUB_TOKEN` for the live lockfile set.
4. Job transpiles every live/placed grammar (may refresh `core.go`).
5. Job runs `go test` on ubuntu on that tree.
6. Job opens a PR as `github-actions[bot]` (`GITHUB_TOKEN`).
7. You click **Approve workflows to run** on the PR.
8. Test matrix runs. Merge when green.

Repo setting: allow GitHub Actions to create and approve pull requests.

Test workflow must use `on: pull_request` (opened / synchronize /
reopened) so the approve banner exists. Also run Test on `push` to
`main`.

A `GITHUB_TOKEN` push does **not** start `on: push` jobs. There is no
button for those.

lewbot / PAT: later. Not Phase 1.

---

## Test workflow

Matrix: linux amd64/arm64, darwin amd64/arm64, windows amd64/arm64.
linux/386 is experimental (`continue-on-error`). Runtime tests only.
If a target breaks, fix portability. Do not bring back per-arch codegen.

| Tree | Must be green |
|------|----------------|
| `main` with `core.go`, no lang packages | Harness + `./grammar` unit tests. Fixture tests **skip** when `grammar.List()` is empty. Do not `t.Fatal` on empty registry. |
| Cook PR after langs land | fixtures for registered langs: no null root, no `HasError` on golden sources |

Fixtures and goldens are the contract. Cook PRs must not invent broken
trees. `LiveParseReady` is optional.

---

## Correctness bar

- Local cook of the live set transpiles and parses.
- After a cook PR merges: same parse on CI (three OS) for registered langs.
- Goldens live in `testdata/` and are human-owned.

---

## Later (not Phase 1)

Uncomment one `#grammar` (or a small batch) → `mod lock` → local prove →
dispatch → cook PR.

First expansion set (corpus fixtures, `testdata/*/sample.*`):

```
ada angular blade COBOL gdscript gitcommit just nim
php php_only purescript rescript squirrel vim wgsl wgsl_bevy
```

Tiny `tiny.*` langs (json, go, python, …) stay commented until you want them.

Optional later: Renovate/mise pins for clang 14 and leaven. GitHub App
for cook PRs so Test runs without the approve click.

---

## Out of scope for the first drop

- Commit original grammar C trees
- Enable the 16 or the full ~190 inventory
- Multi-arch leaven codegen
- lewbot for cook PRs
- History from `ccgo-tree-sitter`
- Git submodules / split grammar repos (Go modules are enough now)

---

## Copy from origin (adapt, no history)

Origin: `modernc-tree-sitter/ccgo-tree-sitter` @ `20260806-try-codegen`.

- `cmd/codegen/*` (leaven path)
- `grammar/api.go`, `query.go`, `parse.go`, `registry.go`, `status.go`,
  `doc.go`, `lineindex.go`, tests
- `grammar/core.go` from a local cook (commit it)
- `testdata/` (all origin fixtures + goldens)
- workspaced: live `#grammar` only; other langs as comments
- `go.mod` `tool` + `replace` for `github.com/lewtec/leaven`
- mise: Go, workspaced, `conda:clang` 14

Adjust fixture tests so missing langs skip. Do not import a language
package in core tests until that package exists on `main`.

---

## Confirmation

Grill closed. Next act: Phase 1 implementation (local harness + `core.go`
on `main`, then dispatch cook of the live set).
