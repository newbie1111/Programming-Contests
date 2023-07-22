package main

import (
	"bufio"
	"errors"
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

func solve(N int, S []string, A []int) []string {
	var (
		min_index = 0
	)

	for i, v := range A {
		if A[min_index] > v {
			min_index = i
		}
	}

	return append(S[min_index:], S[:min_index]...)
}

func main() {
	input.Buffer(make([]byte, 1<<20), 1<<20)
	var (
		N int
		S []string
		A []int
	)

	fmt.Scan(&N)
	for input.Scan() {
		for i, v := range strings.Split(input.Text(), " ") {
			if i == 0 {
				S = append(S, v)
			} else {
				num, _ := strconv.Atoi(v)
				A = append(A, num)
			}
		}
	}

	ans := solve(N, S, A)
	fmt.Println(strings.Join(ans, "\n"))
}
