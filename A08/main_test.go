package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: `5 5
2 0 0 5 1
1 0 3 0 0
0 8 5 0 2
4 1 0 0 6
0 9 2 7 0
2
2 2 4 5
1 1 5 5
`,
			want: `25
56
`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := &strings.Builder{}
			solve(strings.NewReader(tt.input), got)

			if got.String() != tt.want {
				t.Errorf("got: %s, want: %s", got.String(), tt.want)
			}
		})
	}
}
