package main

import (
	"bufio"
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
	Yes    = "Yes"
	No     = "No"
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

func min_int(s []int) int {
	res := MAXINT
	for _, v := range s {
		if res > v {
			res = v
		}
	}

	return res
}

// solves

type Set map[int]bool

func SetKeys(x Set) []int {
	var keys []int

	for k, _ := range x {
		keys = append(keys, k)
	}

	return keys
}

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

func is_upwardly_compatible(x_price, y_price int, x_func, y_func Set) bool {
	check := 0

	if x_price >= y_price {
		check++
	}

	if IsSubset(y_func, x_func) {
		check++
	}

	if x_price > y_price || len(Difference(y_func, x_func)) != 0 {
		check++
	}

	return check == 3
}

func solve(N, M int, P, C []int, F []Set) string {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if i != j && is_upwardly_compatible(P[i], P[j], F[i], F[j]) {
				return Yes
			}
		}
	}
	return No
}

func main() {
	input.Buffer(make([]byte, 1<<20), 1<<20)

	var (
		N int
		M int
		P []int
		C []int
		F []Set
	)

	fmt.Scan(&N, &M)
	for input.Scan() {
		f := make(Set)

		for i, s := range strings.Split(input.Text(), " ") {
			v, _ := strconv.Atoi(s)
			switch {
			case i == 0:
				P = append(P, v)
			case i == 1:
				C = append(C, v)
			default:
				f[v] = true
			}
		}

		F = append(F, f)
	}

	ans := solve(N, M, P, C, F)
	fmt.Println(ans)

}
