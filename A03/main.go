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
	var n, k int
	scanf("%d %d\n", &n, &k)

	ps := make([]int, n)
	for i := 0; i < n; i++ {
		scanf("%d", &ps[i])
	}
	scanf("\n")

	qs := make([]int, n)
	for i := 0; i < n; i++ {
		scanf("%d", &qs[i])
	}
	scanf("\n")

	// process
	result := "No"
	for _, p := range ps {
		for _, q := range qs {
			if p+q == k {
				result = "Yes"
				break
			}
		}
	}

	// output
	printf("%s\n", result)
}
