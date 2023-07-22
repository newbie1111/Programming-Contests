package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(n int, s []string) int {
	set := make(map[string]int)
	for _, v := range s {
		set[v]++
	}
	return len(set)
}

func main() {
	input := bufio.NewScanner(os.Stdin)

	var (
		n int
		s []string
	)

	input.Scan()
	n, _ = strconv.Atoi(input.Text())

	for i := 0; i < n; i++ {
		input.Scan()
		s = append(s, input.Text())
	}
	fmt.Println(solve(n, s))
}
