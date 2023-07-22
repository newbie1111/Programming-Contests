package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// func dump(ty bool, variable ...any) {
// 	if ty {
// 		for _, v := range variable {
// 			fmt.Printf("%v (%v)\n", v, reflect.TypeOf(v))
// 		}
// 	} else {
// 		for _, v := range variable {
// 			fmt.Println(v)
// 		}
// 	}
// }

func solve(N int, A []int) (ans int) {
	ans = math.MaxInt32

	for _, v := range A {
		bit := 0

		for bit = 0; v>>bit&1 != 1; bit++ {
		}

		if ans > bit {
			ans = bit
		}
	}

	return
}

func main() {
	input := bufio.NewScanner(os.Stdin)

	var (
		n int
		s []int
	)

	input.Scan()
	n, _ = strconv.Atoi(input.Text())

	input.Scan()
	for _, v := range strings.Split(input.Text(), " ") {
		a, _ := strconv.Atoi(v)
		s = append(s, a)
	}

	// fmt.Println(n, s)
	// for _, v := range s {
	// 	fmt.Printf("%d ", v)
	// }
	// fmt.Println()
	fmt.Println(solve(n, s))
}
