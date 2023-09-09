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
	var a, b int
	scanf("%d %d\n", &a, &b)

	// process
	result := "No"
	for i := a; i <= b; i++ {
		if 100%i == 0 {
			result = "Yes"
			break
		}
	}

	// output
	printf("%s\n", result)
}
