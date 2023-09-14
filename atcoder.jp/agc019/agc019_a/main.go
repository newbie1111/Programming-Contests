package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

/*
problem solver
*/

func main() {
	input.Split(bufio.ScanWords)

	input.Scan()
	q, _ := strconv.ParseInt(input.Text(), 0, 64)
	input.Scan()
	h, _ := strconv.ParseInt(input.Text(), 0, 64)
	input.Scan()
	s, _ := strconv.ParseInt(input.Text(), 0, 64)
	input.Scan()
	d, _ := strconv.ParseInt(input.Text(), 0, 64)
	input.Scan()
	n, _ := strconv.ParseInt(input.Text(), 0, 64)

	ans := solve(q, h, s, d, n)
	fmt.Println(ans)
}

func solve(q, h, s, d, n int64) any {
	ans, _ := MinValues(8*q, 4*h, 2*s, d)
	ans *= n / 2

	if n%2 != 0 {
		v, _ := MinValues(4*q, 2*h, s)
		ans += v
	}

	return ans
}

/*
initialize common variables
*/

var (
	input             = bufio.NewScanner(os.Stdin)
	output, debug     = log.New(os.Stdout, "", 0), log.New(io.Discard, "DEBUG : ", log.Lshortfile)
	YES, Yes, yes     = "YES", "Yes", "yes"
	NO, No, no        = "NO", "No", "no"
	Even              = "Even"
	Odd               = "Odd"
	mod, mod998244353 = 1e9 + 7, 998244353
)

func init() {
	input.Buffer(make([]byte, 1<<30), 1<<30)

	var (
		verbose bool
	)

	flag.BoolVar(&verbose, "verbose", false, "")
	flag.Parse()

	if verbose {
		debug.SetOutput(os.Stderr)
	}
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
	SignedInteger | UnSignedInteger
}

type SignedInteger interface {
	int | int32 | int64
}
type UnSignedInteger interface {
	uint | uint32 | uint64
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

/*
Data Structures
- Set
- Queue
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

type SliceQueue[T comparable] []T

func NewSliceQueue[T comparable](vars ...T) SliceQueue[T] {
	var (
		newSliceQueue = make(SliceQueue[T], 0, len(vars))
	)

	for _, v := range vars {
		newSliceQueue.Push(v)
	}
	return newSliceQueue
}

func (sq *SliceQueue[T]) Empty() bool {
	return len(*sq) == 0
}

func (sq *SliceQueue[T]) Push(x T) {
	*sq = append(*sq, x)
}

func (sq *SliceQueue[T]) Pop() (T, error) {
	if sq.Empty() {
		return *new(T), errors.New("size of queue is zero")
	} else {
		v := (*sq)[0]
		(*sq)[0] = *new(T) // see also https://github.com/golang/go/wiki/SliceTricks
		*sq = (*sq)[1:]

		return v, nil
	}
}

func (sq *SliceQueue[T]) Rotate(n int) {
	if !sq.Empty() {
		rotateIndex := n % len(*sq)

		if rotateIndex < 0 {
			rotateIndex = len(*sq) + rotateIndex
		}

		*sq = append((*sq)[rotateIndex:], (*sq)[:rotateIndex]...)
	}
}

/*
Character / String
- String Reversal
- string to integer slice
*/

func ReverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func ParseSignedInts[SI SignedInteger](line, sep string) ([]SI, error) {
	var (
		splitWords = strings.Split(line, sep)
		res        = make([]SI, 0, len(splitWords))
	)

	for _, v := range splitWords {
		if n, err := strconv.ParseInt(v, 10, 64); err == nil {
			res = append(res, SI(n))
		} else {
			return make([]SI, 0), errors.New("ParseSingedInts: Parse Error")
		}
	}

	return res, nil
}

func ParseUnSignedInts[USI UnSignedInteger](line, sep string) ([]USI, error) {
	var (
		splitWords = strings.Split(line, sep)
		res        = make([]USI, 0, len(splitWords))
	)

	for _, v := range splitWords {
		if n, err := strconv.ParseUint(v, 10, 64); err == nil {
			res = append(res, USI(n))
		} else {
			return make([]USI, 0), errors.New("ParseUnSingedInts: Parse Error")
		}
	}

	return res, nil
}

/*
Convert value
- Absolute value
- Power
*/

// return absolute integer
func Abs[R Real](x R) R {
	if x >= 0 {
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
- Enumerate all divisors
- check prime number
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

func Divisors[Z Zahl](n Z) []Z {
	div := make([]Z, 0)

	for i := Z(1); i*i < Abs(n); i++ {
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

func IsPrime[Z Zahl](n Z) bool {
	switch {
	case n < 2:
		return false
	case n == 2:
		return true
	case n != 2 && n%2 == 0:
		return false
	default:
		for i := Z(3); i*i <= n; i = i + 2 {
			if n%i == 0 {
				return false
			}
		}
		return true
	}
}

func CumulativeSum[R Real](vars []R) []R {
	var (
		cumsum = make([]R, len(vars)+1)
	)

	if len(vars) != 0 {
		for i, v := range vars {
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
- check all values
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

func CheckAllValues[T comparable](fn func(v T) bool, values ...T) bool {
	res := true

	for _, v := range values {
		res = res && fn(v)
	}

	return res
}

/*
Search Algorithm
- Binary Search
- Measuring Worm Algorithm
*/

func BinarySearch[Z Zahl](negative, positive, dist Z,
	IsContinue func(negative, positive, dist Z) bool,
	IsPositive func(mid Z) bool) (Z, Z) {

	for IsContinue(negative, positive, dist) {
		mid := negative + Abs(positive-negative)/2

		if IsPositive(mid) {
			positive = mid
		} else {
			negative = mid
		}
	}

	return negative, positive
}

func MeasuringWormAlgorithm[Z Zahl](n Z,
	isProgressive func(left, right Z) bool,
	changeRight func(right Z),
	changeLeft func(left Z)) (min, max, interval Z) {
	var (
		left, right Z
	)

	min = n + 1

	for left = Z(0); left < n; left++ {
		for right < n && isProgressive(left, right) {
			changeRight(right)
			right++
		}

		max = Max(max, right-left)
		min = Min(min, right-left)
		interval += Abs(right - left)

		if left == right {
			right++
		} else {
			changeLeft(left)
		}

	}

	return min, max, interval
}
