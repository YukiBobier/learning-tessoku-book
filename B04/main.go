package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// preparation for fast I/O
	reader := bufio.NewReader(os.Stdin)
	scanf := func(f string, a ...interface{}) { _, _ = fmt.Fscanf(reader, f, a...) }

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	printf := func(f string, a ...interface{}) { _, _ = fmt.Fprintf(writer, f, a...) }

	// input
	var n int
	scanf("%b\n", &n)

	// process
	// nothing to do

	// output
	printf("%d\n", n)
}
