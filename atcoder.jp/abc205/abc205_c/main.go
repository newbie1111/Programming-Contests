package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
)

var (
	input  = bufio.NewScanner(os.Stdin)
	YES    = "YES"
	NO     = "NO"
	MAXINT = int(^uint(0) >> 1)
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

func max_int_with_err(s []int) (int, error) {
	var (
		res = MININT
		err = errors.New("length is 0")
	)

	if len(s) != 0 {
		res := s[0]
		err = nil

		for _, v := range s {
			if res < v {
				res = v
			}
		}
	}

	return res, err
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

func min_int_with_err(s []int) (int, error) {
	var (
		res = MAXINT
		err = errors.New("length is 0")
	)

	if len(s) != 0 {
		res := s[0]
		err = nil

		for _, v := range s {
			if res > v {
				res = v
			}
		}
	}

	return res, err
}

func ReverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func abs_i2i(num int) int {
	if num >= 0 {
		return num
	} else {
		return -num
	}
}

func solve(A, B, C int) string {
	var (
		ans    string
		even   = C%2 == 0
		a_plus = A >= 0
		b_plus = B >= 0
	)

	// fmt.Println(A, abs_i2i(A), B, abs_i2i(B))

	if even {
		if abs_i2i(A) < abs_i2i(B) {
			ans = "<"
		} else if abs_i2i(A) > abs_i2i(B) {
			ans = ">"
		} else {
			ans = "="
		}
	} else {
		if a_plus && !b_plus {
			ans = ">"
		} else {
			ans = "<"
		}
	}

	return ans
}

func solve2(A, B, C int) string {
	var (
		ans string
	)

	// fmt.Println(A, abs_i2i(A), B, abs_i2i(B))

	if C%2 == 0 {
		if abs_i2i(A) < abs_i2i(B) {
			ans = "<"
		} else if abs_i2i(A) > abs_i2i(B) {
			ans = ">"
		} else {
			ans = "="
		}
	} else {
		if A < B {
			ans = "<"
		} else if A > B {
			ans = ">"
		} else {
			ans = "="
		}
	}

	return ans
}

func main() {
	input.Buffer(make([]byte, 1<<20), 1<<20)
	var (
		A, B, C int
	)

	fmt.Scan(&A, &B, &C)
	ans := solve2(A, B, C)
	fmt.Println(ans)

}
