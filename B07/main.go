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
	var t int
	_, _ = fmt.Fscanf(input, "%d\n", &t)

	var n int
	_, _ = fmt.Fscanf(input, "%d\n", &n)

	lrs := make([]struct{ l, r int }, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscanf(input, "%d %d\n", &lrs[i].l, &lrs[i].r)
	}

	// process
	differences := make([]int, t)
	for _, lr := range lrs {
		startIdx := lr.l
		endIdx := lr.r
		differences[startIdx]++
		if endIdx < t {
			differences[endIdx]--
		}
	}

	comulativeSumOfDifferences := make([]int, t)
	comulativeSumOfDifferences[0] = differences[0]
	for i := 1; i < t; i++ {
		comulativeSumOfDifferences[i] = comulativeSumOfDifferences[i-1] + differences[i]
	}

	// writing output
	for _, sum := range comulativeSumOfDifferences {
		_, _ = fmt.Fprintf(output, "%d\n", sum)
	}
}
