package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

var (
	input  = bufio.NewScanner(os.Stdin)
	YES    = "YES"
	NO     = "NO"
	MAXINT = int(^uint8(0) >> 1)
	MININT = -MAXINT - 1
)

func dump(variable ...interface{}) {
	for _, v := range variable {
		fmt.Printf("%v : %v\n", reflect.TypeOf(v), v)
	}
}

func byte_to_int(r byte) int {
	return int(r - '0')
}

func max_int(s []int) int {
	res := MININT

	for _, v := range s {
		if res < v {
			res = v
		}
	}

	return res
}

func min_int(s []int) int {
	res := MAXINT
	for _, v := range s {
		if res > v {
			res = v
		}
	}

	return res
}

func ReverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func solve(N int, S []string) int {
	m := make(map[string]bool)

	for _, v := range S {
		if !m[v] && !m[ReverseString(v)] {
			m[v] = true
		}
	}

	return len(m)
}

func main() {
	input.Buffer(make([]byte, 1<<20), 1<<20)
	var (
		N int
		S []string
	)
	fmt.Scan(&N)

	for input.Scan() {
		S = append(S, input.Text())
	}

	ans := solve(N, S)
	fmt.Println(ans)
}
