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
			input: `2
`,
			want: `4
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
