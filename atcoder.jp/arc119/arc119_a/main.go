package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var (
	input    = bufio.NewScanner(os.Stdin)
	YES      = "YES"
	NO       = "NO"
	MAXINT   = int(^uint(0) >> 1)
	MININT   = -MAXINT - 1
	MAXINT64 = int64(^uint64(0) >> 1)
	MININT64 = -MAXINT64 - 1
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
func min_int64(s []int64) int64 {
	res := MAXINT64
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

func ArithmeticProgressionsSum(a0, d, n int) (sum int64, err error) {
	if n < 0 {
		sum, err = 0, errors.New("number of terms must not be negative.")
	} else {
		var (
			a0_i64 = int64(a0)
			d_i64  = int64(d)
			n_i64  = int64(n)
		)
		sum, err = (a0_i64*n_i64)+n_i64*(n_i64-1)/2*d_i64, nil
	}

	return sum, err
}
func GeometicProgressionsSum(a0, r, n int) (sum int64, err error) {
	if n < 0 {
		sum, err = 0, errors.New("number of terms must not be negative.")
	} else {
		var (
			a0_f64 = float64(a0)
			r_f64  = float64(r)
			n_f64  = float64(n)
		)
		sum, err = int64(a0_f64*(math.Pow(r_f64, n_f64)-1)/(r_f64-1)), nil
	}

	return sum, err
}

func GCD(x, y int) int {
	if y == 0 {
		return x
	} else {
		return GCD(y, x%y)
	}
}

func LCM(x, y int) int {
	return x / GCD(x, y) * y
}

func InputListInt(sl *[]int) bool {
	if input.Scan() {
		for _, s := range strings.Split(input.Text(), " ") {
			n, _ := strconv.ParseInt(s, 0, 0)
			*sl = append(*sl, int(n))
		}
		return true
	} else {
		return false
	}
}

func solve(N int64) int64 {
	var (
		abc []int64
	)

	for b := int64(0); b <= 60; b++ {
		pow2b := int64(1 << b)

		if a := N / pow2b; a == 0 {
			break
		} else {
			c := N - a*pow2b
			abc = append(abc, a+b+c)
		}
	}
	return min_int64(abc)
}

func init() {
	input.Buffer(make([]byte, 1<<20), 1<<20)
}

func main() {
	var (
		N int64
	)

	fmt.Scan(&N)

	ans := solve(N)
	fmt.Println(ans)
}
