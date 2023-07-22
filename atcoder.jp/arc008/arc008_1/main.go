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

func solve(N int) int {
	if N%10 <= 6 {
		return 100*(N/10) + 15*(N%10)
	} else {
		return 100 * (N/10 + 1)
	}
}

func main() {
	input.Buffer(make([]byte, 1<<20), 1<<20)
	var n int
	fmt.Scan(&n)

	ans := solve(n)
	fmt.Println(ans)
}
