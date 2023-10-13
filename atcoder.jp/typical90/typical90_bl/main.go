package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
problem solver
*/

func main() {
	var (
		N, Q       int
		A, L, R, V []int
	)

	fmt.Scan(&N, &Q)
	input.Scan()
	A, _ = ParseSignedInts[int](input.Text(), " ")
	for input.Scan() {
		line, _ := ParseSignedInts[int](input.Text(), " ")
		L = append(L, line[0])
		R = append(R, line[1])
		V = append(V, line[2])
	}

	debug.Println(
		N, Q,
		A, L, R, V,
	)
	ans := solve(
		N, Q,
		A, L, R, V,
	)
	output.Println(ans)
}

func solve(
	N, Q int,
	A, L, R, V []int,
) any {
	var (
		E          = make([]int, N-1)
		vecL, vecR = NewVector[int](L...), NewVector[int](R...)
		sum        int
		ans        []int
	)

	vecL.ScalarSubtract(1)
	vecR.ScalarSubtract(1)

	for i := 0; i < N-1; i++ {
		E[i] = A[i+1] - A[i]
		sum += Abs(E[i])
	}

	debug.Println(E)

	for i := 0; i < Q; i++ {
		left, right := vecL[i]-1, vecR[i]+1

		debug.Println(left, right)

		if left >= 0 {
			E[left] += V[i]
			sum += V[i]
		}

		if right < N-1 {
			E[right] += V[i]
			sum += V[i]
		}

		debug.Println(E)
		ans = append(ans, sum)
	}

	return FormatSignedInts(ans, 10, "\n")
}

/*
initialize common variables
*/

var (
	input         = bufio.NewScanner(os.Stdin)
	output, debug = log.New(os.Stdout, "", 0), log.New(io.Discard, "DEBUG : ", log.Lshortfile)
)

const (
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
	int | int8 | int16 | int32 | int64
}

type UnSignedInteger interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type Float interface {
	float32 | float64
}

/*
Data Structures
- Set
- Queue
- Circular Queue
- Vector
- Matrix
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

type CircularQueue[Any any] struct {
	head, tail int
	cq         []Any
}

func NewCircularQueue[Any any]() CircularQueue[Any] {
	newCircularQueue := *new(CircularQueue[Any])
	newCircularQueue.InIt()
	return newCircularQueue
}

func (cq *CircularQueue[Any]) InIt() {
	cq.head, cq.tail = 0, 0
	cq.cq = make([]Any, 1)
}

func (cq *CircularQueue[Any]) nextIndex(index int) int {
	return (index + 1) % len(cq.cq)

}

func (cq *CircularQueue[Any]) previousIndex(index int) int {
	return (len(cq.cq) + index - 1) % len(cq.cq)

}

func (cq *CircularQueue[Any]) Full() bool {
	return cq.head == cq.nextIndex(cq.tail)
}

func (cq *CircularQueue[Any]) Empty() bool {
	return cq.head == cq.tail
}

func (cq *CircularQueue[Any]) Length() int {
	return Abs(cq.tail - cq.head)
}

func (cq *CircularQueue[Any]) toSlice() []Any {
	res := make([]Any, 0, cq.Length())
	for i := cq.head; i != cq.tail; i = cq.nextIndex(i) {
		res = append(res, cq.cq[i])
	}
	return res
}

func (cq *CircularQueue[Any]) PushLeft(x Any) {
	if cq.Full() {
		prevCQ := cq.toSlice()
		newCQ := make([]Any, len(cq.cq)<<1)

		newCQ[0] = x
		copy(newCQ[1:], prevCQ)

		cq.head, cq.tail = 0, len(prevCQ)+1
		cq.cq = newCQ
	} else {
		cq.head = cq.previousIndex(cq.head)
		cq.cq[cq.head] = x
	}
}

func (cq *CircularQueue[Any]) PushRight(x Any) {
	if cq.Full() {
		prevCQ := cq.toSlice()
		newCQ := make([]Any, len(cq.cq)<<1)

		copy(newCQ, prevCQ)
		newCQ[len(prevCQ)] = x

		cq.head, cq.tail = 0, len(prevCQ)+1
		cq.cq = newCQ
	} else {
		cq.cq[cq.tail] = x
		cq.tail = cq.nextIndex(cq.tail)
	}
}

func (cq *CircularQueue[Any]) PopRight() Any {
	if cq.Empty() {
		return *new(Any)
	} else {
		cq.tail = cq.previousIndex(cq.tail)
		res := cq.cq[cq.tail]
		return res
	}
}

func (cq *CircularQueue[Any]) At(index int) Any {
	return cq.cq[(cq.head+index)%len(cq.cq)]
}

func (cq *CircularQueue[Any]) SetAt(index int, value Any) {
	cq.cq[(cq.head+index)%len(cq.cq)] = value
}

type Vector[R Real] []R

func NewVector[R Real](v ...R) Vector[R] {
	newVec := new(Vector[R]).InIt(len(v))
	copy(newVec, v)
	return newVec
}

func (v *Vector[R]) InIt(dim int) Vector[R] {
	return make(Vector[R], dim)
}

func (v *Vector[R]) Reshape(dim int) Vector[R] {
	nv := new(Vector[R]).InIt(dim)
	copy(nv, *v)
	return nv
}

func (v *Vector[R]) Dimension() int {
	return len(*v)
}

func (v *Vector[R]) ScalarAdd(n R) {
	for i := 0; i < v.Dimension(); i++ {
		(*v)[i] += n
	}
}

func (v *Vector[R]) ScalarSubtract(n R) {
	for i := 0; i < v.Dimension(); i++ {
		(*v)[i] -= n
	}
}

func (v *Vector[R]) Difference(diff Vector[R]) Vector[R] {
	nv := NewVector[R](*v...)

	if v.Dimension() != diff.Dimension() {
		nv.Reshape(Max(v.Dimension(), diff.Dimension()))
	}

	for index, value := range diff {
		nv[index] -= value
	}

	return nv
}

func (v *Vector[R]) ScalarMultipul(n R) {
	for i := 0; i < v.Dimension(); i++ {
		(*v)[i] *= n
	}
}

func (v *Vector[R]) ScalarDivide(n R) {
	for i := 0; i < v.Dimension(); i++ {
		(*v)[i] /= n
	}
}

type Matrix[R Real] []Vector[R]

func NewMatrix[R Real](rows ...[]R) Matrix[R] {
	maxDim := 0
	for _, row := range rows {
		maxDim = Max(maxDim, len(row))
	}

	newMat := new(Matrix[R]).InIt(len(rows), maxDim)
	for i, row := range rows {
		copy(newMat[i], row)
	}

	return newMat
}

func (mat *Matrix[R]) InIt(row, col int) Matrix[R] {
	newMat := make(Matrix[R], 0, row*col)
	for i := 0; i < row; i++ {
		rowVector := NewVector[R]()
		rowVector.Reshape(col)
		newMat = append(newMat, rowVector)
	}
	return newMat
}

func (mat *Matrix[R]) Size() int {
	if len(*mat) == 0 {
		return 0
	} else {
		return len(*mat) * len((*mat)[0])
	}
}

func (mat *Matrix[R]) Dimension() (int, int) {
	if mat.Size() == 0 {
		return 0, 0
	} else {
		return len(*mat), len((*mat)[0])
	}
}

func (mat *Matrix[R]) RowAt(index int) Vector[R] {
	row := (*mat)[index][:]
	return NewVector(row...)
}

func (mat *Matrix[R]) ColAt(index int) Vector[R] {
	var (
		r, _ = mat.Dimension()
		col  = new(Vector[R]).InIt(r)
	)

	for i, vec := range *mat {
		col[i] = vec[index]
	}

	return col
}

/*
Character / String
- String Reversal
- string to integer slice
- integer slice to string
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

func FormatSignedInts[SI SignedInteger](s []SI, base int, sep string) string {
	var (
		res []string
	)

	for _, v := range s {
		res = append(res, strconv.FormatInt(int64(v), base))
	}

	return strings.Join(res, sep)
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
- Ceil
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

func Ceil[Z Zahl](a, b Z) Z {
	return (a + b - 1) / b
}

/*
Number theoretic algorithm
- total
- factorial
- product
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

func Factorial[Z Zahl](n Z) Z {
	res := make([]Z, 0)
	for i := Z(1); i < n; i++ {
		res = append(res, i)
	}
	return Product[Z](res...)
}

func Product[R Real](vars ...R) R {
	if len(vars) == 0 {
		return 0
	} else {
		product := R(1)
		for _, v := range vars {
			product *= v
		}
		return product
	}
}

func ProductWithMod[Z Zahl](mod Z, vars ...Z) Z {
	if len(vars) == 0 {
		return 0
	} else {
		product := Z(1)
		for _, v := range vars {
			product *= v
			product %= mod
		}
		return product
	}
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
		sum, err = 0, errors.New("number of terms must not be negative")
	} else {
		nReal := R(n)
		sum, err = (a0*nReal)+nReal*(nReal-1)/2*d, nil
	}

	return sum, err
}

func GeometicProgressionsSum[Z Zahl](a0, r, n Z) (sum Z, err error) {
	if n < 0 {
		sum, err = 0, errors.New("number of terms must not be negative")
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

func BinarySearch[R Real](negative, positive, dist R,
	IsContinue func(negative, positive, dist R) bool,
	IsPositive func(mid R) bool) (R, R) {

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
