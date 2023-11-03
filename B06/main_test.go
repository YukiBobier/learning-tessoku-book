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
			input: `7
0 1 1 0 1 0 0
3
2 5
2 7
5 7
`,
			want: `win
draw
lose
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
