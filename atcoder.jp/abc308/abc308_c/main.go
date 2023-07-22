package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

var (
	input = bufio.NewScanner(os.Stdin)
	YES   = "YES"
	NO    = "NO"
)

func dump(variable ...interface{}) {
	for _, v := range variable {
		fmt.Printf("%v : %v\n", reflect.TypeOf(v), v)
	}
}

func solve(N int, A, B []int) []string {

	type tuple struct {
		numerator   int // 分子
		denominator int // 分母
		index       int // 人i
	}

	var (
		person []tuple
		ans    []string
	)

	for i := 0; i < N; i++ {
		person = append(person, tuple{-A[i], A[i] + B[i], i + 1})
	}

	sort.SliceStable(person, func(i, j int) bool {
		s := person[i].numerator*person[j].denominator - person[i].denominator*person[j].numerator
		return s < 0
	})

	// dump(person)

	for _, v := range person {
		s := strconv.Itoa(v.index)
		ans = append(ans, s)
	}

	return ans

}

func main() {
	var (
		n int
		a []int
		b []int
	)

	input.Scan()
	n, _ = strconv.Atoi(input.Text())

	for i := 0; i < n; i++ {
		input.Scan()
		v := strings.Split(input.Text(), " ")
		av, _ := strconv.Atoi(v[0])
		bv, _ := strconv.Atoi(v[1])
		a = append(a, av)
		b = append(b, bv)
	}

	ans := solve(n, a, b)
	fmt.Println(strings.Join(ans, " "))
}
