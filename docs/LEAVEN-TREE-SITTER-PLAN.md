# Plan: `lewtec/leaven-tree-sitter`

Port of the leaven-based tree-sitter Go stack from the
`ccgo-tree-sitter` experiment (`20260806-try-codegen` branch) into a
**fresh** repo. Same product scheme as `ccgo-tree-sitter`, backend is
**leaven only** (no ccgo / modernc.org/libc).

Source of truth for this plan: grilling session 2026-08-07/08.

---

## Goals

1. **Well-tested small set** of languages for short CI iteration.
2. **Eventually** enable the full workspaced inventory after polish.
3. Commit **generated Go**, not original grammar **C**.
4. C sources: **workspaced JIT** download → transpile when regenerating.

---

## Product shape

| Item | Decision |
|------|----------|
| Target repo | https://github.com/lewtec/leaven-tree-sitter |
| Port style | **Clean port** (no history from `ccgo-tree-sitter`) |
| Module path | `github.com/lewtec/leaven-tree-sitter` |
| Backend | clang 14 → LLVM IR → `go tool leaven` (lewtec/leaven pin) |
| Module layout | **One module for core** (`…/grammar`) + **one module per language** (`…/grammar/<lang>`) |
| Committed | Generated `.go` (`core.go`, `grammar/<lang>/grammar.go`, hand API, fixtures, codegen, lock) |
| Not committed | Original C under `third-party/tree-sitter-*` (JIT via workspaced) |
| Scheme | Codegen → nested modules → `Register` → parse/fixtures API (like ccgo-tree-sitter) |

---

## Enabled languages (day one)

**Full inventory** declared in workspaced; **most commented out**.

**Enabled** = richer **corpus-stolen** fixtures only (`testdata/*/sample.*`):

```
ada angular blade COBOL gdscript gitcommit just nim
php php_only purescript rescript squirrel vim wgsl wgsl_bevy
```

- First **local canary:** `php`
- Tiny `tiny.*` fixture langs (json, go, python, …) stay **commented** until later
- Later: uncomment workspaced entry → cook → PR → merge

Approx generated size for the 16: **~110–120 MB** Go (COBOL is the whale ~31 MB).

---

## Landing process

### Phase 1 — Setup (push to `main`)

Until the factory works, land directly on **`main`**:

- Module root + `go.mod` / nested language module wiring
- Hand API: `api.go`, `query.go`, `parse.go`, `registry.go`, `status.go`, …
- `cmd/codegen` (leaven path only)
- workspaced: full grammar list, only the 16 enabled (rest commented)
- `workspaced.lock.json` pins
- **Fixture sources + goldens** for the 16 (human-owned contract)
- CI: **tests every commit** on ubuntu + macos + windows
- CI: **codegen workflow** (see below), not yet cooking langs until smoke is green

### Phase 2 — Cook (codegen workflow)

1. Local prove: `codegen` for **core + php** (JIT C → leaven → tests).
2. Push factory to `main` when local is green.
3. **workflow_dispatch** and/or **lockfile change** triggers codegen.
4. Codegen runs **once on ubuntu** (clang 14 + leaven), not multi-arch codegen.
5. Bot opens a **PR** with generated `.go` (`GITHUB_TOKEN` / `github-actions[bot]`).
6. **Tests run on that PR** (multi-OS).
7. Merge when green.

### Phase 3 — Expand

Uncomment more langs in workspaced → same cook path. No need for 190 on day one.

---

## CI detail

| Workflow | Trigger | What |
|----------|---------|------|
| **Test** | Every commit / PR | `go test` for enabled packages + fixtures; matrix **ubuntu + macos + windows** |
| **Codegen** | `workflow_dispatch`, lockfile (and/or schedule if desired) | JIT place C → transpile **enabled** set → open PR with diffs |

- Codegen does **not** re-run multi-GOOS/GOARCH matrices (leaven = single pure-Go output).
- Multi-OS is for **runtime tests**; if something breaks, fix portability—do not revive platform-split codegen.
- Public repo: free GHA minutes; prefer fast iteration via **small enabled set**.

---

## Correctness bar

- Enabled langs must **transpile** and **parse** with the hand API the way current fixtures do (no null root, no `HasError` on golden sources).
- Goldens live in setup (`testdata/`); cook PRs should not invent broken trees.
- `LiveParseReady`-style probe optional; fixtures are the contract.

---

## Local iteration (fast path)

1. Work on factory **locally** (seconds–minutes for `--only=php`).
2. Do **not** debug first clang/PATH issues only on GHA.
3. When local canary is green → push main → dispatch cook → PR → multi-OS tests.
4. Then widen the enabled set.

---

## What is *not* in scope for the first drop

- Committing original grammar C trees
- Enabling all ~85 tiny-fixture langs or all 190
- Multi-arch leaven codegen matrix
- lewbot for codegen PRs (use Actions bot first)
- History transplant from `ccgo-tree-sitter`

---

## Source material (this monorepo)

Useful origin branch: `20260806-try-codegen` on `modernc-tree-sitter/ccgo-tree-sitter`.

Copy / adapt (not full history):

- `cmd/codegen/*` (leaven)
- `grammar/api.go`, `query.go`, `parse.go`, `registry.go`, `status.go`, `doc.go`, `lineindex.go`, tests
- `grammar/core.go` (or produce via first cook)
- Enabled `grammar/<lang>/` after cook
- `testdata/` for the 16
- workspaced grammar declarations (comment pattern)
- leaven pin in `go.mod` (`github.com/lewtec/leaven` replace)

---

## Open / later

- When to uncomment more langs (after polish)
- Optional: drop COBOL from enable set to shrink first cook PR
- Optional: add `json` as a tiny seed for README without full corpus set
- Renovate/mise pins for clang 14 and leaven

---

## Confirmation

Grill decisions closed as of this document. Implementation starts when we act on Phase 1 (local factory + push main).
