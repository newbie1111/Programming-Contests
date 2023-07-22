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

func solve(n int, a []int) string {
	sl := make([]string, n)

	for i, v := range a {
		sl[v-1] = strconv.Itoa(i + 1)
	}

	return strings.Join(sl, " ")
}

func main() {
	var (
		n int
	)

	input.Buffer(make([]byte, 1<<20), 1<<20)
	input.Split(bufio.ScanWords)

	input.Scan()
	n, _ = strconv.Atoi(input.Text())
	a := make([]int, n)

	for i := 0; input.Scan(); i++ {
		val, _ := strconv.Atoi(input.Text())
		a[i] = val
	}

	ans := solve(n, a)
	fmt.Println(ans)
}
