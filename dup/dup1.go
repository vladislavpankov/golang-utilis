package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	count := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		count[input.Text()]++
	}

	for line, n := range count {
		if n > 1 {
			fmt.Printf("number of lines:%d\t contain:%s\n", n, line)
		}
	}
}
