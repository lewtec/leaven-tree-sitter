package main

// fixtureSnippet is one source file under testdata/<language>/<File>.
type fixtureSnippet struct {
	File   string
	Source string
}

// fixtureSources: one minimal, valid sample per generated grammar language.
// Goldens are testdata/<lang>/<File>.golden.json (UPDATE_GOLDENS=1 to refresh).
var fixtureSources = map[string]fixtureSnippet{
	"arduino": {
		File:   "tiny.ino",
		Source: "void setup() {}\nvoid loop() {}\n",
	},
	"astro": {
		File:   "tiny.astro",
		Source: "---\nconst x = 1\n---\n<p>{x}</p>\n",
	},
	"bash": {
		File:   "tiny.sh",
		Source: "echo hello\n",
	},
	"beancount": {
		File:   "tiny.beancount",
		Source: "2020-01-01 open Assets:Cash\n",
	},
	"bicep": {
		File:   "tiny.bicep",
		Source: "param name string = 'world'\n",
	},
	"c": {
		File:   "tiny.c",
		Source: "int main(void) { return 0; }\n",
	},
	"c_sharp": {
		File:   "tiny.cs",
		Source: "class C { static void Main() {} }\n",
	},
	"cairo": {
		File:   "tiny.cairo",
		Source: "fn main() {}\n",
	},
	"cmake": {
		File:   "tiny.cmake",
		Source: "project(example)\n",
	},
	"commonlisp": {
		File:   "tiny.lisp",
		Source: "(defun hello () 1)\n",
	},
	"core_schema": {
		File:   "tiny.schema",
		Source: "true",
	},
	"cpp": {
		File:   "tiny.cpp",
		Source: "int main() { return 0; }\n",
	},
	"css": {
		File:   "tiny.css",
		Source: "a { color: red; }\n",
	},
	"csv": {
		File:   "tiny.csv",
		Source: "a,b,c\n1,2,3\n",
	},
	"cuda": {
		File:   "tiny.cu",
		Source: "__global__ void k() {}\n",
	},
	"dart": {
		File:   "tiny.dart",
		Source: "void main() {}\n",
	},
	"diff": {
		File:   "tiny.diff",
		Source: "--- a/x\n+++ b/x\n@@ -1 +1 @@\n-old\n+new\n",
	},
	"dockerfile": {
		File:   "tiny.Dockerfile",
		Source: "FROM alpine\n",
	},
	"dtd": {
		File:   "tiny.dtd",
		Source: "<!ELEMENT note (#PCDATA)>\n",
	},
	"elixir": {
		File:   "tiny.ex",
		Source: "defmodule M do\n  def f, do: 1\nend\n",
	},
	"gitattributes": {
		File:   "tiny.gitattributes",
		Source: "*.go text\n",
	},
	"glsl": {
		File:   "tiny.glsl",
		Source: "void main() {}\n",
	},
	"go": {
		File:   "tiny.go",
		Source: "package main\n\nfunc main() {}\n",
	},
	"hare": {
		File:   "tiny.ha",
		Source: "export fn main() void = void;\n",
	},
	"haskell": {
		File:   "tiny.hs",
		Source: "main = return ()\n",
	},
	"hcl": {
		File:   "tiny.hcl",
		Source: "resource \"null\" \"x\" {}\n",
	},
	"html": {
		File:   "tiny.html",
		Source: "<html><body></body></html>\n",
	},
	"java": {
		File:   "tiny.java",
		Source: "class Main { public static void main(String[] a) {} }\n",
	},
	"javascript": {
		File:   "tiny.js",
		Source: "const x = 1;\n",
	},
	"jsdoc": {
		File:   "tiny.jsdoc",
		Source: "/**\n * @param {string} x\n */\n",
	},
	"json": {
		File:   "tiny.json",
		Source: "{\"key\": [1, true]}\n",
	},
	"json_schema": {
		File:   "tiny.schema.json",
		Source: "true",
	},
	"julia": {
		File:   "tiny.jl",
		Source: "x = 1\n",
	},
	"kotlin": {
		File:   "tiny.kt",
		Source: "fun main() {}\n",
	},
	"legacy_schema": {
		File:   "tiny.schema",
		Source: "true",
	},
	"lua": {
		File:   "tiny.lua",
		Source: "local x = 10\nprint(x)\n",
	},
	"luadoc": {
		File:   "tiny.luadoc",
		Source: "@class Foo\n",
	},
	"luap": {
		File:   "tiny.luap",
		Source: "a*b\n",
	},
	"make": {
		File:   "tiny.mk",
		Source: "all:\n\t@echo hi\n",
	},
	"markdown": {
		File:   "tiny.md",
		Source: "# Title\n\nHello.\n",
	},
	"markdown_inline": {
		File:   "tiny.inline.md",
		Source: "**bold** and `code`\n",
	},
	"nix": {
		File:   "tiny.nix",
		Source: "{ x = 1; }\n",
	},
	"ocaml": {
		File:   "tiny.ml",
		Source: "let x = 1\n",
	},
	"ocaml_interface": {
		File:   "tiny.mli",
		Source: "val x : int\n",
	},
	"ocaml_type": {
		File:   "tiny.mly",
		Source: "int\n",
	},
	"po": {
		File:   "tiny.po",
		Source: "msgid \"hi\"\nmsgstr \"hello\"\n",
	},
	"psv": {
		File:   "tiny.psv",
		Source: "a|b|c\n1|2|3\n",
	},
	"puppet": {
		File:   "tiny.pp",
		Source: "notify { 'hi': }\n",
	},
	"python": {
		File:   "tiny.py",
		Source: "x = 10\nprint(x)\n",
	},
	"query": {
		File:   "tiny.scm",
		Source: "(identifier) @name\n",
	},
	"regex": {
		File:   "tiny.regex",
		Source: "a+b*\n",
	},
	"requirements": {
		File:   "tiny.requirements.txt",
		Source: "requests==2.0.0\n",
	},
	"ron": {
		File:   "tiny.ron",
		Source: "(a: 1)\n",
	},
	"ruby": {
		File:   "tiny.rb",
		Source: "puts 1\n",
	},
	"rust": {
		File:   "tiny.rs",
		Source: "fn main() {}\n",
	},
	"scala": {
		File:   "tiny.scala",
		Source: "object Main { def main(args: Array[String]) = () }\n",
	},
	"solidity": {
		File:   "tiny.sol",
		Source: "contract C {}\n",
	},
	"svelte": {
		File:   "tiny.svelte",
		Source: "<script>let x = 1</script>\n<p>{x}</p>\n",
	},
	"tcl": {
		File:   "tiny.tcl",
		Source: "puts hello\n",
	},
	"templ": {
		File:   "tiny.templ",
		Source: "package main\n\ntempl Hello() {\n\t<div>hi</div>\n}\n",
	},
	"terraform": {
		File:   "tiny.tf",
		Source: "resource \"null_resource\" \"x\" {}\n",
	},
	"toml": {
		File:   "tiny.toml",
		Source: "x = 1\n",
	},
	"tsv": {
		File:   "tiny.tsv",
		Source: "a\tb\tc\n1\t2\t3\n",
	},
	"tsx": {
		File:   "tiny.tsx",
		Source: "export const x = <div />;\n",
	},
	"typescript": {
		File:   "tiny.ts",
		Source: "const x: number = 1;\n",
	},
	"vue": {
		File:   "tiny.vue",
		Source: "<template><p>hi</p></template>\n",
	},
	"xml": {
		File:   "tiny.xml",
		Source: "<?xml version=\"1.0\"?><a/>\n",
	},
	"yaml": {
		File:   "tiny.yaml",
		Source: "x: 1\n",
	},
	"zig": {
		File:   "tiny.zig",
		Source: "pub fn main() void {}\n",
	},
}
