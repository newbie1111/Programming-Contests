package main

import (
	"bufio"
	"fmt"
	"os"
)

func price(yen500, yen100, yen50 int) int {
	return yen500*500 + yen100*100 + yen50*50
}

func solve(a, b, c, x int) (ans int) {

	ans = 0
	for i := 0; i < a+1; i++ {
		for j := 0; j < b+1; j++ {
			for k := 0; k < c+1; k++ {
				if price(i, j, k) == x {
					ans++
				}
			}
		}
	}
	return
}

func main() {
	input := bufio.NewReader(os.Stdin)

	var a, b, c, x int

	fmt.Fscan(input, &a, &b, &c, &x)

	ans := solve(a, b, c, x)
	fmt.Println(ans)
}
