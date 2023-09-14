package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

/*
problem solver
*/

func solve[Z Zahl](n Z) any {
	switch x := Sum(Divisors(n)...) - n; {
	case x < n:
		return "Deficient"
	case x == n:
		return "Perfect"
	default:
		return "Abundant"
	}
}

func Divisors[Z Zahl](n Z) []Z {
	div := make([]Z, 0)

	for i := Z(1); i*i < AbsInteger(n); i++ {
		if n%i == 0 {
			div = append(div, i)
			if n/i != i {
				div = append(div, n/i)
			}
		}
	}

	if n < 0 {
		divCopy := make([]Z, len(div))
		copy(divCopy, div)
		for _, v := range divCopy {
			div = append(div, -v)
		}
	}

	return div
}

func main() {
	var n int64
	fmt.Scan(&n)
	ans := solve(n)
	fmt.Println(ans)
}

/*
initialize common variables
*/

var (
	input             = bufio.NewScanner(os.Stdin)
	output, debug     = log.New(os.Stdout, "", 0), log.New(os.Stderr, "DEBUG : ", log.Lshortfile)
	YES, Yes, yes     = "YES", "Yes", "yes"
	NO, No, no        = "NO", "No", "no"
	Even              = "Even"
	Odd               = "Odd"
	mod, mod998244353 = 1e9 + 7, 998244353
)

func init() {
	input.Buffer(make([]byte, 1<<30), 1<<30)
}

/*
Generics Type Interface
- Number
	- Integer
	- Float
*/

type Real interface {
	Zahl | Float
}

type Zahl interface {
	SignedInteger | uint | uint32 | uint64
}

type SignedInteger interface {
	int | int32 | int64
}

type Float interface {
	float32 | float64
}

/*
Input / Output
*/

func dump(variable ...interface{}) {
	for _, v := range variable {
		debug.Printf("%v : %v\n", reflect.TypeOf(v), v)
	}
}

func InputListSignedInteger[SI SignedInteger](sep string) ([]SI, error) {
	var (
		res []SI
	)

	for _, s := range strings.Split(input.Text(), sep) {
		n, err := strconv.ParseInt(s, 10, 64)

		if err != nil {
			return []SI{}, err
		}

		res = append(res, SI(n))
	}

	return res, nil
}

/*
Data Structures
- Set
*/

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](vars ...T) Set[T] {
	var (
		newSet = make(Set[T])
	)

	for _, v := range vars {
		newSet.Add(v)
	}

	return newSet
}

func (s *Set[T]) Add(x T) {
	(*s)[x] = struct{}{}
}

func (s *Set[T]) Remove(x T) {
	delete(*s, x)
}

func (s *Set[T]) Exist(x T) bool {
	_, exist := (*s)[x]
	return exist
}

func (s *Set[T]) Empty() bool {
	return len(*s) == 0
}

/*
Character / String
- String Reversal
*/

func ReverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

/*
Convert value
- Absolute value
- Power
*/

// return absolute integer
func AbsInteger[Z Zahl](x Z) Z {
	if x > 0 {
		return x
	} else {
		return -x
	}
}

func PowInteger[Z Zahl](a, n Z) Z {
	var (
		res = Z(1)
	)

	for n > 0 {
		if n&1 == 1 {
			res *= a
		}
		a, n = a*a, n>>1
	}

	return res
}

/*
Number theoretic algorithm
- total
- Cumulative Sum
- Arithmetic Progressions
- Geometic Progressions
- Greatest Common Divisor
- Least Common Multipul
*/

func Sum[R Real](vars ...R) R {
	var (
		sum R
	)

	for _, v := range vars {
		sum += v
	}

	return sum
}

func CumulativeSum[R Real](vars []R) []R {
	var (
		cumsum = make([]R, len(vars))
	)

	if len(vars) != 0 {
		cumsum[0] = vars[0]
		for i, v := range vars[1:] {
			cumsum[i+1] = v + cumsum[i]
		}
	}

	return cumsum
}

func ArithmeticProgressionsSum[R Real, Z Zahl](a0, d R, n Z) (sum R, err error) {
	if n < 0 {
		sum, err = 0, errors.New("number of terms must not be negative.")
	} else {
		nReal := R(n)
		sum, err = (a0*nReal)+nReal*(nReal-1)/2*d, nil
	}

	return sum, err
}

func GeometicProgressionsSum[Z Zahl](a0, r, n Z) (sum Z, err error) {
	if n < 0 {
		sum, err = 0, errors.New("number of terms must not be negative.")
	} else {
		sum, err = a0*(PowInteger(r, n)-1)/(r-1), nil
	}

	return sum, err
}

func GreatestCommonDivisor[Z Zahl](x, y Z) Z {
	if y == 0 {
		return x
	} else {
		return GreatestCommonDivisor(y, x%y)
	}
}

func GreatestCommonDivisorIntegers[Z Zahl](vars ...Z) (gcd Z, err error) {
	if len(vars) == 0 {
		return 0, errors.New("vars length is 0")
	} else {
		gcd, err = vars[0], nil

		for _, v := range vars[1:] {
			gcd = GreatestCommonDivisor(gcd, v)
		}

		return gcd, err
	}
}

func LeastCommonMultipul[Z Zahl](x, y Z) Z {
	return x / GreatestCommonDivisor(x, y) * y
}

func LeastCommonMultipulIntegers[I Zahl](vars ...I) (lcm I, err error) {
	if len(vars) == 0 {
		return 0, errors.New("vars length is 0")
	} else {
		lcm, err = vars[0], nil

		for _, v := range vars[1:] {
			lcm = LeastCommonMultipul(lcm, v)
		}

		return lcm, nil
	}
}

/*
Comparison of Values
- larger / smaller
- largest / smallest
*/

func Max[R Real](x, y R) R {
	if x >= y {
		return x
	} else {
		return y
	}
}

func Min[R Real](x, y R) R {
	if x >= y {
		return y
	} else {
		return x
	}
}

func MaxValues[R Real](vars ...R) (R, error) {
	var (
		res R
		err error
	)

	if len(vars) == 0 {
		err = errors.New("length is 0")
	} else {
		res = vars[0]
		for _, v := range vars[1:] {
			res = Max(res, v)
		}
	}

	return res, err
}

func MinValues[R Real](vars ...R) (R, error) {
	var (
		res R
		err error
	)

	if len(vars) == 0 {
		err = errors.New("length is 0")
	} else {
		res = vars[0]
		for _, v := range vars[1:] {
			res = Min(res, v)
		}
	}

	return res, err
}

/*
Search Algorithm
- Binary Search
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
