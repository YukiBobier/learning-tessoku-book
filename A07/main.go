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
	var d int
	_, _ = fmt.Fscanf(input, "%d\n", &d)

	var n int
	_, _ = fmt.Fscanf(input, "%d\n", &n)

	lrs := make([]struct{ l, r int }, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscanf(input, "%d %d\n", &lrs[i].l, &lrs[i].r)
	}

	// process
	sumOfAttendees := make([]int, d)
	for _, lr := range lrs {
		startIdx := lr.l - 1
		endIdx := lr.r - 1
		for i := startIdx; i <= endIdx; i++ {
			sumOfAttendees[i]++
		}
	}

	// writing output
	for _, sum := range sumOfAttendees {
		_, _ = fmt.Fprintf(output, "%d\n", sum)
	}
}
