package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(s string) int {
	ans := 0
	for _, c := range s {
		if c == '1' {
			ans++
		}
	}
	return ans
}

func main() {
	input := bufio.NewReader(os.Stdin)

	var (
		s string
	)

	fmt.Fscan(input, &s)

	ans := solve(s)
	fmt.Println(ans)
}
