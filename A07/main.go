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

	// process
	answer := n * n

	// writing output
	_, _ = fmt.Fprintf(output, "%d\n", answer)
}
