# leaven-tree-sitter

Tree-sitter in pure Go via clang → LLVM IR → [leaven](https://github.com/lewtec/leaven). No CGO, no ccgo.

Uncomment a `#grammar` in `workspaced.cue`, lock, dispatch codegen. The factory cooks whatever is live.

```bash
mise install
export CC="$(mise which clang)"
mise run grammars:lock
mise run grammars:sync          # JIT C under third-party/ (gitignored)
mise run codegen                # core + enabled langs; discard grammar/<lang>/ before commit
mise run test                   # or: go run ./cmd/codegen --only=ada
```

See `docs/LEAVEN-TREE-SITTER-PLAN.md`.
