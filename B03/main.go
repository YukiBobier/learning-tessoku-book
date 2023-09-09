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
	scanf("%d\n", &n)

	as := make([]int, n)
	for i := 0; i < n; i++ {
		scanf("%d", &as[i])
	}
	scanf("\n")

	// process
	result := "No"
	for i := range as {
		for j := range as {
			for k := range as {
				if i == j || j == k || i == k {
					continue
				}

				if as[i]+as[j]+as[k] == 1000 {
					result = "Yes"
					break
				}
			}
		}
	}

	// output
	printf("%s\n", result)
}
