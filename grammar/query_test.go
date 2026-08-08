package grammar

import "testing"

func TestExecuteMatchesNilGuards(t *testing.T) {
	// Non-nil query with garbage pointer: root checks must prevent native use.
	fakeQ := &Query{q: new(TSQuery)}

	cases := []struct {
		name string
		q    *Query
		root *Node
	}{
		{"nil query", nil, nil},
		{"zero query", &Query{}, nil},
		{"zero query zero node", &Query{}, &Node{}},
		{"unusable query nil root", fakeQ, nil},
		{"unusable query zero node", fakeQ, &Node{}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.q.ExecuteMatches(tc.root, nil)
			if len(got) != 0 {
				t.Fatalf("len(matches)=%d want 0", len(got))
			}
		})
	}
}
