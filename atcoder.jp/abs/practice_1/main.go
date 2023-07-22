package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(a, b, c int, s string) (int, string) {
	return a + b + c, s
}

func main() {
	input := bufio.NewReader(os.Stdin)

	var (
		a int
		b int
		c int
		s string
	)

	fmt.Fscan(input, &a, &b, &c, &s)

	abc, s := solve(a, b, c, s)
	fmt.Println(abc, s)
}
