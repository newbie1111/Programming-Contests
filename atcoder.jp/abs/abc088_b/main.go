package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func solve(N int, A []int) int {
	alice, bob := 0, 0
	sort.Slice(A, func(i int, j int) bool { return A[i] > A[j] })

	for i := 0; i < len(A); i++ {
		if i%2 == 0 {
			alice += A[i]
		} else {
			bob += A[i]
		}
	}

	return alice - bob
}

func main() {
	input := bufio.NewScanner(os.Stdin)

	var (
		N int
		A []int
	)

	input.Scan()
	N, _ = strconv.Atoi(input.Text())

	input.Scan()
	for _, v := range strings.Split(input.Text(), " ") {
		a, _ := strconv.Atoi(v)
		A = append(A, a)
	}

	// fmt.Println(n, s)
	// for _, v := range s {
	// 	fmt.Printf("%d ", v)
	// }
	// fmt.Println()
	fmt.Println(solve(N, A))
}
