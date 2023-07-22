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

func solve(P []int) string {
	var (
		ans []string
	)

	for _, n := range P {
		ans = append(ans, string(rune('a')+int32(n-1)))
	}

	return strings.Join(ans, "")
}

func main() {
	input.Buffer(make([]byte, 1<<20), 1<<20)
	var (
		P []int
	)

	input.Scan()
	for _, s := range strings.Split(input.Text(), " ") {
		num, _ := strconv.Atoi(s)
		P = append(P, num)
	}

	ans := solve(P)
	fmt.Println(ans)
}
