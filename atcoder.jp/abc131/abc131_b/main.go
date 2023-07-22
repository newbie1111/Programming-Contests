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

func solve(N, L int) int {
	ans := (2*L + N - 1) * N / 2
	if (L+1-1) <= 0 && 0 <= (L+N-1) {
		return ans
	} else if (L + N - 1) < 0 {
		return ans - (L + N - 1)
	} else {
		return ans - (L + 1 - 1)
	}
}

func main() {
	var N, L int
	fmt.Scan(&N, &L)
	ans := solve(N, L)
	fmt.Println(ans)
}
