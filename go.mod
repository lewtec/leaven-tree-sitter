module github.com/lewtec/leaven-tree-sitter

go 1.25.0

require (
	github.com/spf13/cobra v1.10.2
	golang.org/x/mod v0.38.0
)

require github.com/google/go-cmp v0.6.0

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/lewtec/leaven v0.0.0-20260808132831-ddad4fe451d9 // indirect
	github.com/lewtec/leaven-tree-sitter/grammar v0.0.0
	github.com/llir/ll v0.0.0-20210719001141-246f2b6b1fa9 // indirect
	github.com/llir/llvm v0.3.5-0.20220120022543-ad080cea55b7 // indirect
	github.com/mewmew/float v0.0.0-20201204173432-505706aa38fa // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/spf13/pflag v1.0.9 // indirect
	golang.org/x/exp v0.0.0-20220407100705-7b9b53b0aca4 // indirect
	golang.org/x/sync v0.21.0 // indirect
	golang.org/x/tools v0.47.0 // indirect
)

tool github.com/lewtec/leaven/cmd/leaven

replace github.com/lewtec/leaven-tree-sitter/grammar => ./grammar
