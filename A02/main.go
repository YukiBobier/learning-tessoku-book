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
	var n, x int
	scanf("%d %d\n", &n, &x)

	as := make([]int, n)
	for i := 0; i < n; i++ {
		scanf("%d", &as[i])
	}

	// process
	result := "No"
	for _, a := range as {
		if a == x {
			result = "Yes"
			break
		}
	}

	// output
	printf("%s\n", result)
}
