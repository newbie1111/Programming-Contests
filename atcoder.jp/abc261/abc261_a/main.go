package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
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

func max_int(s []int) int {
	res := -9223372036854775808

	for _, v := range s {
		if res < v {
			res = v
		}
	}

	return res
}
func min_int(s []int) int {
	res := 9223372036854775807

	for _, v := range s {
		if res > v {
			res = v
		}
	}

	return res
}

func solve(left, right []int) int {
	return max_int([]int{min_int(right) - max_int(left), 0})
}

func main() {
	var (
		left  []int
		right []int
	)

	input.Scan()
	for i, v := range strings.Split(input.Text(), " ") {
		val, _ := strconv.Atoi(v)
		if i%2 == 0 {
			left = append(left, val)
		} else {
			right = append(right, val)
		}
	}

	ans := solve(left, right)
	fmt.Println(ans)
}
