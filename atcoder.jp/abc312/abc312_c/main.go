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
	input            = bufio.NewScanner(os.Stdin)
	YES, Yes, yes    = "YES", "Yes", "yes"
	NO, No, no       = "NO", "No", "no"
	Even             = "Even"
	Odd              = "Odd"
	MAXINT, MAXINT64 = int(^uint(0) >> 1), int64(^uint64(0) >> 1)
	MININT, MININT64 = -MAXINT - 1, -MAXINT64 - 1
)

// Set

type Set map[int]bool

func (st *Set) NewSet(vars ...int) *Set {
	var s = make(Set)
	for _, v := range vars {
		s[v] = true
	}
	return &s
}

func (st *Set) SetKeys() []int {
	var keys []int

	for k, _ := range *st {
		keys = append(keys, k)
	}

	return keys
}

func (st *Set) Add(v int) { (*st)[v] = true }

func (st *Set) Exist(v int) bool { return (*st)[v] }

func (st *Set) Empty() bool { return len(*st) == 0 }

func (st *Set) Union(y Set) Set {
	var res = make(Set)

	for k, _ := range *st {
		res[k] = true
	}

	for k, _ := range y {
		res[k] = true
	}

	return res
}

func (st *Set) Difference(y Set) Set {
	var res = make(Set)

	for k, _ := range *st {
		if !y[k] {
			res[k] = true
		}
	}

	return res
}

func (st *Set) Intersection(y Set) Set {
	var res = make(Set)

	for k, _ := range *st {
		if y[k] {
			res[k] = true
		}
	}

	return res
}

func (st *Set) IsEqual(y Set) bool {
	var (
		x_keys            = st.SetKeys()
		y_keys            = y.SetKeys()
		intersection_xy   = st.Intersection(y)
		intersection_keys = intersection_xy.SetKeys()
	)

	return len(x_keys) == len(intersection_keys) && len(y_keys) == len(intersection_keys)
}

func (st *Set) IsSubset(x, y Set) bool { return len(y.Difference(x)) == 0 }

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

func solve(N, M int, A, B []int) interface{} {
	var (
		negative, positive = 0, int(math.Pow10(10))
		buyers, sellers    = make([]int, M), make([]int, N)
	)

	copy(sellers, A)
	copy(buyers, B)

	for positive-negative > 1 {
		var (
			mid = negative + (positive-negative)>>1
		)

		isOk := func() bool {
			var (
				sell, buy int
			)

			for _, v := range sellers {
				if v <= mid {
					sell++
				}
			}

			for _, v := range buyers {
				if v >= mid {
					buy++
				}
			}

			return sell >= buy
		}

		if isOk() {
			positive = mid
		} else {
			negative = mid
		}
	}

	return positive
}

func BinarySearch(negative, positive, dist interface{},
	IsContinue func(negative, positive, dist interface{}) bool,
	HowToMiddle func(negative, positive interface{}) interface{},
	IsPositive func(mid interface{}) bool,
	IsReturnPositive bool) interface{} {

	for IsContinue(negative, positive, dist) {
		mid := HowToMiddle(negative, positive)

		if IsPositive(mid) {
			positive = mid
		} else {
			negative = mid
		}
	}

	if IsReturnPositive {
		return positive
	} else {
		return negative
	}
}

func solve2(N, M int, A, B []int) interface{} {
	var (
		negative, positive = 0, int(1e10)
	)

	return BinarySearch(negative, positive, 1,
		func(negative, positive, dist interface{}) bool {
			return positive.(int)-negative.(int) > dist.(int)
		},
		func(negative, positive interface{}) interface{} {
			return negative.(int) + (positive.(int)-negative.(int))>>1
		},
		func(mid interface{}) bool {
			var (
				sell, buy int
			)

			for _, v := range A {
				if mid.(int) >= v {
					sell++
				}
			}

			for _, v := range B {
				if mid.(int) <= v {
					buy++
				}
			}

			return sell >= buy
		}, true).(int)

}

func init() {
	input.Buffer(make([]byte, 1<<30), 1<<30)
}

func main() {
	var (
		N, M int
		A, B []int
	)

	fmt.Scan(&N, &M)
	input.Scan()
	A, _ = InputListInt()
	input.Scan()
	B, _ = InputListInt()

	ans := solve2(N, M, A, B)
	fmt.Println(ans)
}
