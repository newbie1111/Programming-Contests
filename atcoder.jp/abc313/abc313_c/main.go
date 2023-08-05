package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

/*
problem solver
*/

func solve(N int, A []int) interface{} {
	var (
		sum, _ = SumInt(A...)
		B      = make([]int, len(A))
		ans    int
	)

	sort.Slice(A, func(i, j int) bool { return A[i] < A[j] })

	for i := 0; i < N; i++ {
		B[i] = sum / N
	}

	for i := 0; i < sum%N; i++ {
		B[N-1-i]++
	}

	for i := 0; i < N; i++ {
		ans += AbsInt(A[i] - B[i])
	}

	return ans / 2
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func AbsInt64(x int64) int64 {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func main() {
	var (
		N int
		A []int
	)

	fmt.Scan(&N)
	input.Scan()
	A, _ = InputListInt()
	ans := solve(N, A)
	fmt.Println(ans)
}

/*
initialize common variables
*/

var (
	input            = bufio.NewScanner(os.Stdin)
	YES, Yes, yes    = "YES", "Yes", "yes"
	NO, No, no       = "NO", "No", "no"
	Even             = "Even"
	Odd              = "Odd"
	MAXINT, MAXINT64 = int(^uint(0) >> 1), int64(^uint64(0) >> 1)
	MININT, MININT64 = -MAXINT - 1, -MAXINT64 - 1
)

func init() {
	input.Buffer(make([]byte, 1<<30), 1<<30)
}

/*
Input / Output
*/

func dump(variable ...interface{}) {
	for _, v := range variable {
		fmt.Printf("%v : %v\n", reflect.TypeOf(v), v)
	}
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

/*
Data Structures
*/

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

/*
Character / String
*/

func byte_to_int(r byte) int {
	return int(r - '0')
}

func ReverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

/*
Number theoretic algorithm
*/

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

func GreatestCommonDivisorInt(x, y int) int {
	if y == 0 {
		return x
	} else {
		return GreatestCommonDivisorInt(y, x%y)
	}
}

func GreatestCommonDivisorIntegers(vars ...int) (gcd int, err error) {
	if len(vars) == 0 {
		return -1, errors.New("vars length is 0")
	} else {
		gcd, err = vars[0], nil

		for i := 1; i < len(vars); i++ {
			gcd = GreatestCommonDivisorInt(gcd, vars[i])
		}

		return gcd, err
	}
}

func LeastCommonMultipulInt(x, y int) int {
	return x / GreatestCommonDivisorInt(x, y) * y
}

func LeastCommonMultipulIntegers(vars ...int) (lcm int, err error) {
	if len(vars) == 0 {
		return -1, errors.New("vars length is 0")
	} else {
		lcm, err = vars[0], nil

		for i := 1; i < len(vars); i++ {
			lcm = LeastCommonMultipulInt(lcm, vars[i])
		}

		return lcm, nil
	}
}

/*
Comparison of Values
*/

func MaxInt(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

func MaxIntegers(vars ...int) (int, error) {
	var (
		res int
		err error
	)

	if len(vars) == 0 {
		res, err = MAXINT, errors.New("vars length is 0")
	} else {
		res, err = vars[0], nil

		for _, v := range vars {
			res = MaxInt(res, v)
		}
	}

	return res, err
}

func MinInt(x, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}

func MinIntegers(vars ...int) (int, error) {
	var (
		res int
		err error
	)

	if len(vars) == 0 {
		res, err = MININT, errors.New("vars length is 0")
	} else {
		res, err = vars[0], nil

		for _, v := range vars {
			res = MinInt(res, v)
		}
	}

	return res, err
}

func MaxFloat64(x, y float64) float64 {
	if x >= y {
		return x
	} else {
		return y
	}
}

func MinFloat64(x, y float64) float64 {
	if x <= y {
		return x
	} else {
		return y
	}
}

func MaxInt64(x, y int64) int64 {
	if x >= y {
		return x
	} else {
		return y
	}
}

func MaxInteger64s(vars ...int64) (int64, error) {
	var (
		res int64
		err error
	)

	if len(vars) == 0 {
		res, err = MAXINT64, errors.New("vars length is 0")
	} else {
		res, err = vars[0], nil

		for _, v := range vars {
			res = MaxInt64(res, v)
		}
	}

	return res, err
}

func MinInt64(x, y int64) int64 {
	if x <= y {
		return x
	} else {
		return y
	}
}

func MinInteger64s(vars ...int64) (int64, error) {
	var (
		res int64
		err error
	)

	if len(vars) == 0 {
		res, err = MININT64, errors.New("vars length is 0")
	} else {
		res, err = vars[0], nil

		for _, v := range vars {
			res = MinInt64(res, v)
		}
	}

	return res, err
}

/*
Search Algorithm
*/

func BinarySearch(negative, positive, dist interface{},
	IsContinue func(negative, positive, dist interface{}) bool,
	HowToMiddle func(negative, positive interface{}) interface{},
	IsPositive func(mid interface{}) bool) (interface{}, interface{}) {

	for IsContinue(negative, positive, dist) {
		mid := HowToMiddle(negative, positive)

		if IsPositive(mid) {
			positive = mid
		} else {
			negative = mid
		}
	}

	return negative, positive
}

func SumInt(vars ...int) (sum int, err error) {
	if len(vars) == 0 {
		err = errors.New("length is 0")
	} else {
		for _, v := range vars {
			sum += v
		}
	}
	return
}
