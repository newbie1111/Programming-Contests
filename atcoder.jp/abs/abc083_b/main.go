package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(n, a, b int) (ans int) {
	ans = 0

	for i := 1; i <= n; i++ {
		target := 0
		
		for _, v := range strings.Split(strconv.Itoa(i), "") {
			num, _ := strconv.Atoi(v)
			target += num
		}

		if a <= target && target <= b {
			ans += i
		}
	}
	return
}

func main() {
	input := bufio.NewReader(os.Stdin)

	var n, a, b int

	fmt.Fscan(input, &n, &a, &b)

	ans := solve(n, a, b)
	fmt.Println(ans)
}
