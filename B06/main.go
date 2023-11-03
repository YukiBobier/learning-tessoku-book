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
	var n int
	_, _ = fmt.Fscanf(input, "%d\n", &n)

	as := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscanf(input, "%d", &as[i])
	}
	_, _ = fmt.Fscanf(input, "\n")

	var q int
	_, _ = fmt.Fscanf(input, "%d\n", &q)

	lrs := make([]struct{ l, r int }, q)
	for i := 0; i < q; i++ {
		_, _ = fmt.Fscanf(input, "%d %d\n", &lrs[i].l, &lrs[i].r)
	}

	// process
	comulativeSumOfHits := make([]int, n)
	comulativeSumOfHits[0] = as[0]
	for i := 1; i < n; i++ {
		comulativeSumOfHits[i] = comulativeSumOfHits[i-1] + as[i]
	}

	answers := make([]string, q)
	for i := 0; i < q; i++ {
		startIdx := lrs[i].l - 1
		endIdx := lrs[i].r - 1

		var sumOfHit int
		if startIdx == 0 {
			sumOfHit = comulativeSumOfHits[endIdx]
		} else {
			sumOfHit = comulativeSumOfHits[endIdx] - comulativeSumOfHits[startIdx-1]
		}

		length := endIdx - startIdx + 1
		threshold := length / 2
		if length%2 == 0 {
			switch {
			case sumOfHit < threshold:
				answers[i] = "lose"
			case sumOfHit == threshold:
				answers[i] = "draw"
			case sumOfHit > threshold:
				answers[i] = "win"
			}
		} else {
			switch {
			case sumOfHit <= threshold:
				answers[i] = "lose"
			case sumOfHit > threshold:
				answers[i] = "win"
			}
		}
	}

	// writing output
	for _, answer := range answers {
		_, _ = fmt.Fprintf(output, "%s\n", answer)
	}
}
