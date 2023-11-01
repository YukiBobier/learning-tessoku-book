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
			input: "15 3\n62 65 41 13 20 11 18 44 53 12 18 17 14 10 39\n4 13\n3 10\n2 15\n",
			want:  "220\n212\n375\n",
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
