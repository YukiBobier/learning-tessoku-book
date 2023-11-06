package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	stdout := bufio.NewWriter(os.Stdout)
	defer stdout.Flush()

	solve(stdin, stdout)
}

func solve(input io.Reader, output io.Writer) {
	// reading input
	var h, w int
	_, _ = fmt.Fscanf(input, "%d %d\n", &h, &w)

	x := make([][]int, h)
	for i := 0; i < h; i++ {
		x[i] = make([]int, w)
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			_, _ = fmt.Fscanf(input, "%d", &x[i][j])
		}
		_, _ = fmt.Fscanf(input, "\n")
	}

	var q int
	_, _ = fmt.Fscanf(input, "%d\n", &q)

	abcds := make([]struct{ a, b, c, d int }, q)
	for i := 0; i < q; i++ {
		_, _ = fmt.Fscanf(input, "%d %d %d %d\n", &abcds[i].a, &abcds[i].b, &abcds[i].c, &abcds[i].d)
	}

	// process
	comulativeSums := make([][]int, h)
	for i := 0; i < h; i++ {
		comulativeSums[i] = make([]int, w)
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			comulativeSums[i][j] = x[i][j]
			if i > 0 {
				comulativeSums[i][j] += comulativeSums[i-1][j]
			}
			if j > 0 {
				comulativeSums[i][j] += comulativeSums[i][j-1]
			}
			if i > 0 && j > 0 {
				comulativeSums[i][j] -= comulativeSums[i-1][j-1]
			}
		}
	}

	answers := make([]int, q)
	for i := 0; i < q; i++ {
		aIdx, bIdx, cIdx, dIdx := abcds[i].a-1, abcds[i].b-1, abcds[i].c-1, abcds[i].d-1
		answers[i] = comulativeSums[cIdx][dIdx]
		if aIdx > 0 {
			answers[i] -= comulativeSums[aIdx-1][dIdx]
		}
		if bIdx > 0 {
			answers[i] -= comulativeSums[cIdx][bIdx-1]
		}
		if aIdx > 0 && bIdx > 0 {
			answers[i] += comulativeSums[aIdx-1][bIdx-1]
		}
	}

	// writing output
	for _, ans := range answers {
		_, _ = fmt.Fprintf(output, "%d\n", ans)
	}
}
