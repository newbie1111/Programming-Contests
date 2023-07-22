package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(a, b int) string {
	if a&b&1 == 1 {
		return "Odd"
	} else {
		return "Even"
	}
}

func main() {
	input := bufio.NewReader(os.Stdin)

	var (
		a int
		b int
	)

	fmt.Fscan(input, &a, &b)

	ans := solve(a, b)
	fmt.Println(ans)
}
