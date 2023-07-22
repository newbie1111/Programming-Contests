package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
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

func clockwize(y, x, left, right, top, bottom int) (int, int) {
	dx, dy := 0, 0

	if y == top {
		if x == (right - 1) {
			dx = 1
		} else {
			dy = 1
		}
	} else if y == (bottom - 1) {
		if x == left {
			dx = -1
		} else {
			dy = -1
		}
	}

	if x == left {
		if !(y == top) {
			dx = -1
		}
	} else if x == (right - 1) {
		if !(y == (bottom - 1)) {
			dx = 1
		}
	}

	return y + dx, x + dy

}

func solve(n int, sl []string) [][]byte {
	var ans [][]byte

	for i := 0; i < n; i++ {
		ans = append(ans, make([]byte, n))
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			x, y := clockwize(i, j, 0, n, 0, n)
			// fmt.Println(i, j, x, y)
			ans[x][y] = sl[i][j]
		}
	}

	return ans
}

func main() {
	var (
		n  int
		sl []string
	)

	input.Scan()
	n, _ = strconv.Atoi(input.Text())

	for i := 0; i < n; i++ {
		input.Scan()
		sl = append(sl, input.Text())
	}
	ans := solve(n, sl)
	for _, v := range ans {
		for _, c := range v {
			fmt.Print(string(c))
		}
		fmt.Println()
	}
}
