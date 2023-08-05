package main

import (
	"bufio"
	"errors"
	"math"
	"math/rand"
	"os"
	"testing"
	"time"
)

// testing : go test template_test.go -v
// benchmark : go test template_test.go -v -bench .

// initialize common variables

var (
	input            = bufio.NewScanner(os.Stdin)
	YES, Yes, yes    = "YES", "Yes", "yes"
	NO, No, no       = "NO", "No", "no"
	Even             = "Even"
	Odd              = "Odd"
	MAXINT, MAXINT64 = int(^uint(0) >> 1), int64(^uint64(0) >> 1)
	MININT, MININT64 = -MAXINT - 1, -MAXINT64 - 1
)

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

func TestBinarySearch(t *testing.T) {
	var (
		rand = rand.New(rand.NewSource(time.Now().UnixNano()))
		XInt = rand.Int()
	)

	n, p := BinarySearch(0, MAXINT, 1,
		func(negative, positive, dist interface{}) bool {
			return positive.(int)-negative.(int) > dist.(int)
		},
		func(negative, positive interface{}) interface{} {
			return negative.(int) + (positive.(int)-negative.(int))>>1
		},
		func(mid interface{}) bool {
			return mid.(int) >= XInt
		})

	if !(n.(int) == XInt-1 && p.(int) == XInt) {
		t.Errorf("negative = %d | positive = %d | target = %d", n, p, XInt)
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var (
			rand = rand.New(rand.NewSource(time.Now().UnixNano()))
			XInt = rand.Int()
		)

		BinarySearch(MININT+1, MAXINT, 1,
			func(negative, positive, dist interface{}) bool {
				return positive.(int)-negative.(int) > dist.(int)
			},
			func(negative, positive interface{}) interface{} {
				return negative.(int) + (positive.(int)-negative.(int))>>1
			},
			func(mid interface{}) bool {
				return mid.(int) >= XInt
			})
	}
}

func GreatestCommonDivisorInt(x, y int) int {
	if y == 0 {
		return x
	} else {
		return GreatestCommonDivisorInt(y, x%y)
	}
}

func GreatestCommonDivisorSliceInt(vars ...int) (gcd int, err error) {
	if len(vars) == 0 {
		return -1, errors.New("vars length is 0")
	} else if len(vars) == 1 {
		return vars[0], nil
	} else {
		gcd = vars[0]

		for i := 1; i < len(vars); i++ {
			gcd = GreatestCommonDivisorInt(gcd, vars[i])
		}

		return gcd, nil
	}
}

func TestGreatestCommonDivisorSliceInt(t *testing.T) {
	var (
		sl = []int{6, 12, 18}
	)

	gcd, _ := GreatestCommonDivisorSliceInt(sl...)

	if !(gcd == 6) {
		t.Errorf("target %v | gcd = %d", sl, gcd)
	}
}

func LeastCommonMultipulInt(x, y int) int {
	return x / GreatestCommonDivisorInt(x, y) * y
}

func LeastCommonMultipulSliceInt(vars ...int) (lcm int, err error) {
	if len(vars) == 0 {
		return -1, errors.New("vars length is 0")
	} else if len(vars) == 1 {
		return vars[0], nil
	} else {
		lcm = vars[0]

		for i := 1; i < len(vars); i++ {
			lcm = LeastCommonMultipulInt(lcm, vars[i])
		}

		return lcm, nil
	}
}
func TestLeastCommonMultipulSliceInt(t *testing.T) {
	var (
		sl = []int{6, 12, 18}
	)

	lcm, _ := LeastCommonMultipulSliceInt(sl...)

	if !(lcm == 36) {
		t.Errorf("target %v | lcm = %d", sl, lcm)
	}
}

func PowInt(a, n int) int {
	return int(math.Pow(float64(a), float64(n)))
}
