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
	var n, q int
	_, _ = fmt.Fscanf(input, "%d %d\n", &n, &q)

	as := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscanf(input, "%d", &as[i])
	}
	_, _ = fmt.Fscanf(input, "\n")

	lrs := make([]struct{ l, r int }, q)
	for i := 0; i < q; i++ {
		_, _ = fmt.Fscanf(input, "%d %d\n", &lrs[i].l, &lrs[i].r)
	}

	// process
	sums := make([]int, n)
	sums[0] = as[0]
	for i := 1; i < n; i++ {
		sums[i] = sums[i-1] + as[i]
	}

	answers := make([]int, q)
	for i := 0; i < q; i++ {
		startIdx := lrs[i].l - 1
		endIdx := lrs[i].r - 1
		if startIdx == 0 {
			answers[i] = sums[endIdx]
		} else {
			answers[i] = sums[endIdx] - sums[startIdx-1]
		}
	}

	// writing output
	for _, answer := range answers {
		_, _ = fmt.Fprintf(output, "%d\n", answer)
	}
}
