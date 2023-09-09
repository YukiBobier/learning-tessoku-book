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
	var n, k int
	_, _ = fmt.Fscanf(input, "%d %d\n", &n, &k)

	// process
	counter := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if k-i-j >= 1 && k-i-j <= n {
				counter++
			}
		}
	}

	// writing output
	_, _ = fmt.Fprintf(output, "%d\n", counter)
}
