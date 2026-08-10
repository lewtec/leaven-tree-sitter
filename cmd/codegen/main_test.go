package main

import "testing"

func TestAllGrammarsFailed(t *testing.T) {
	t.Parallel()
	tests := []struct {
		failed, total int
		want          bool
	}{
		{0, 0, false},
		{0, 10, false},
		{1, 191, false},
		{5, 191, false},
		{191, 191, true},
		{1, 1, true},
	}
	for _, tc := range tests {
		got := allGrammarsFailed(tc.failed, tc.total)
		if got != tc.want {
			t.Fatalf("allGrammarsFailed(%d, %d)=%v want %v", tc.failed, tc.total, got, tc.want)
		}
	}
}
