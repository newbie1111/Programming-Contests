package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
)

var (
	input = bufio.NewScanner(os.Stdin)
	YES   = "YES"
	NO    = "NO"
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
	res := math.MinInt

	for _, v := range s {
		if res < v {
			res = v
		}
	}

	return res
}

func min_int(s []int) int {
	res := math.MaxInt

	for _, v := range s {
		if res > v {
			res = v
		}
	}

	return res
}

func solve() {
}

func main() {

}
