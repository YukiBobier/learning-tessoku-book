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
			input: `10
7
0 3
2 4
1 3
0 3
5 6
5 6
5 6
`,
			want: `2
3
4
1
0
3
0
0
0
0
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
