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
	MAXINT = int(^uint8(0) >> 1)
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

// solve
func solve(N, T, M int, A, B []int) int {
	m := make(map[int]Set)

	for i := 0; i < M; i++ {
		if _, ok := m[A[i]]; !ok {
			m[A[i]] = NewSet()
		}
		if _, ok := m[B[i]]; !ok {
			m[B[i]] = NewSet()
		}

		Add(m[A[i]], B[i])
		Add(m[B[i]], A[i])
	}

	res := dfs(T, m)

	return res
}

func dfs(T int, m map[int]Set) int {
	return 0
}

func main() {
	input.Buffer(make([]byte, 1<<20), 1<<20)
	var (
		N, T, M int
		A, B    []int
	)

	fmt.Scan(&N, &T, &M)
	for input.Scan() {
		for i, s := range strings.Split(input.Text(), " ") {
			v, _ := strconv.Atoi(s)
			switch i {
			case 0:
				A = append(A, v)
			case 1:
				B = append(B, v)
			}
		}
	}

	ans := solve(N, T, M, A, B)
	fmt.Println(ans)
}
