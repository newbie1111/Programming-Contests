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

func solve(a, b, k int) (int, int) {
	takahashi, aoki := a, b

	for i := 0; i < k; i++ {
		if i&1 == 0 {
			takahashi >>= 1
			aoki += takahashi
		} else {
			aoki >>= 1
			takahashi += aoki
		}
	}

	return takahashi, aoki
}

func main() {
	var a, b, k int
	fmt.Scan(&a, &b, &k)
	takahashi, aoki := solve(a, b, k)
	fmt.Println(takahashi, aoki)
}
