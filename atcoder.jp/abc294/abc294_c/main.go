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

// Set

type Set map[int]bool

func NewSet(vars ...int) Set {
	var s = make(Set)
	for _, v := range vars {
		s[v] = true
	}
	return s
}

func SetKeys(x Set) []int {
	var keys []int

	for k, _ := range x {
		keys = append(keys, k)
	}

	return keys
}

func Add(x Set, v int) { x[v] = true }

func Exist(x Set, v int) bool { return x[v] }

func Empty(x Set) bool { return len(x) == 0 }

func Union(x, y Set) Set {
	var res = make(Set)

	for k, _ := range x {
		res[k] = true
	}

	for k, _ := range y {
		res[k] = true
	}

	return res
}

func Difference(x, y Set) Set {
	var res = make(Set)

	for k, _ := range x {
		if !y[k] {
			res[k] = true
		}
	}

	return res
}

func Intersection(x, y Set) Set {
	var res = make(Set)

	for k, _ := range x {
		if y[k] {
			res[k] = true
		}
	}

	return res
}

func IsEqual(x, y Set) bool {
	var (
		x_keys            = SetKeys(x)
		y_keys            = SetKeys(y)
		intersection_keys = SetKeys(Intersection(x, y))
	)

	return len(x_keys) == len(intersection_keys) && len(y_keys) == len(intersection_keys)
}

func IsSubset(x, y Set) bool { return len(Difference(y, x)) == 0 }

func solve(N, M int, A, B []int) (CA, CB []string) {
	var (
		ValueSetA = NewSet(A...)
		ValueSetB = NewSet(B...)
		c         = append(A, B...)
	)

	sort.Slice(c, func(i, j int) bool { return c[i] < c[j] })

	for i, v := range c {
		if Exist(ValueSetA, v) {
			CA = append(CA, strconv.Itoa(i+1))
		} else if Exist(ValueSetB, v) {
			CB = append(CB, strconv.Itoa(i+1))
		} else {
			panic("Error")
		}
	}

	// fmt.Println(ValueSetA, ValueSetB, c)
	return
}
func solve2(N, M int, A, B []int) (CA, CB []string) {
	var (
		c         = append(A, B...)
		c_indexes = make(map[int]string)
	)

	sort.Slice(c, func(i, j int) bool { return c[i] < c[j] })

	for i, v := range c {
		c_indexes[v] = strconv.Itoa(i + 1)
	}

	for _, v := range A {
		CA = append(CA, c_indexes[v])
	}

	for _, v := range B {
		CB = append(CB, c_indexes[v])
	}

	return
}

func init() {
	input.Buffer(make([]byte, 1<<20), 1<<20)
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

// func InputListInt() (sl *[]int, err error) {
// 	if input.Scan() {
// 		for _, s := range strings.Split(input.Text(), " ") {
// 			n, _ := strconv.ParseInt(s, 0, 0)
// 			*sl = append(*sl, int(n))
// 		}

// 		return sl, nil
// 	} else {
// 		return sl, errors.New("cant scan")
// 	}
// }

func main() {
	var (
		N, M int
		A, B []int
	)

	fmt.Scan(&N, &M)
	InputListInt(&A)
	InputListInt(&B)
	// ca, cb := solve(N, M, A, B)
	ca, cb := solve2(N, M, A, B)

	fmt.Println(strings.Join(ca, " "))
	fmt.Println(strings.Join(cb, " "))
}
