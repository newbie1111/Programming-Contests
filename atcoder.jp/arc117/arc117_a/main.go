package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
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

func byte_to_int(r byte) int {
	return int(r - '0')
}

func solve(a, b int) []int {
	var (
		plus  []int
		minus []int
		sum   int
	)

	for i := 1; i < a+1; i++ {
		plus = append(plus, i)
	}

	for i := 1; i < b+1; i++ {
		minus = append(minus, -i)
	}

	if a < b {
		for i := len(plus); i < len(minus); i++ {
			sum += minus[i]
		}
		plus[a-1] += -sum
	} else {
		for i := len(minus); i < len(plus); i++ {
			sum += plus[i]
		}
		minus[b-1] += -sum
	}

	// fmt.Println(plus, minus)

	return append(plus, minus...)
}

func main() {
	var (
		a int
		b int
	)

	fmt.Scan(&a, &b)
	ans := solve(a, b)
	for i, v := range ans {
		fmt.Print(v)

		if (i - 1) != len(ans) {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
