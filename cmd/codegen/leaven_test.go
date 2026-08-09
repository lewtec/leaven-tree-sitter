package main

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestRunLeavenTinyIR(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	ll := filepath.Join(dir, "add.ll")
	mustWrite(t, ll, `
define i32 @add(i32 %a, i32 %b) {
  %c = add i32 %a, %b
  ret i32 %c
}
`)
	got, err := new(LeavenTranspiler).runLeaven(t.Context(), ll, "grammar_add")
	if err != nil {
		t.Fatal(err)
	}
	src := string(got)
	if !strings.Contains(src, "package grammar_add") {
		t.Fatalf("missing package clause:\n%s", src)
	}
	if !strings.Contains(src, "func add(") {
		t.Fatalf("missing add func:\n%s", src)
	}
}

func TestRunLeavenInvalidPackage(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	ll := filepath.Join(dir, "add.ll")
	mustWrite(t, ll, `
define i32 @add(i32 %a, i32 %b) {
  %c = add i32 %a, %b
  ret i32 %c
}
`)
	_, err := new(LeavenTranspiler).runLeaven(t.Context(), ll, "123bad")
	if err == nil {
		t.Fatal("expected error")
	}
}
