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

func SumInt(s []int) (sum int) {
	for _, v := range s {
		sum += v
	}
	return
}

func CumsumInt(s []int) (cumsum []int) {
	for i, v := range s {
		if i == 0 {
			cumsum = append(cumsum, v)
		} else {
			cumsum = append(cumsum, v+cumsum[len(cumsum)-1])
		}
	}
	return
}

func Sum_i64(s []int64) (sum int64) {
	for _, v := range s {
		sum += v
	}
	return
}

func Cumsum_i64(s []int64) (cumsum []int64) {
	for i, v := range s {
		if i == 0 {
			cumsum = append(cumsum, v)
		} else {
			cumsum = append(cumsum, v+cumsum[len(cumsum)-1])
		}
	}
	return
}

func InputListInt() ([]int, error) {
	var (
		res []int
	)

	for _, s := range strings.Split(input.Text(), " ") {
		n, err := strconv.ParseInt(s, 0, 0)

		if err != nil {
			return []int{}, err
		}

		res = append(res, int(n))
	}

	return res, nil
}

func solve(N int, B []int) int {
	var (
		A = make([]int, N)
	)

	A[0], A[len(A)-1] = B[0], B[len(B)-1]
	for i := 1; i < N-1; i++ {
		A[i] = min_int(B[i-1 : i+1])
	}

	// fmt.Println(N, A, B)
	return SumInt(A)
}

func init() {
	input.Buffer(make([]byte, 1<<20), 1<<20)
}

func main() {
	var (
		N int
	)

	fmt.Scan(&N)

	input.Scan()
	B, _ := InputListInt()
	ans := solve(N, B)
	fmt.Println(ans)
}
