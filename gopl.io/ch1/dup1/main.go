package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	// 注意： 忽略input.Erro() 中可能的错误
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t\n", n, line)
		}
	}
}
