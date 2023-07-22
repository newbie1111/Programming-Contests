package main

import (
	"bufio"
	"fmt"
	"os"
)

func price(yen10000, yen5000, yen1000 int) int {
	return yen10000*10000 + yen5000*5000 + yen1000*1000
}

func solve(N, Y int) (x, y, z int) {
	x, y, z = -1, -1, -1

	for yen10k := 0; yen10k <= N; yen10k++ {
		for yen5k := 0; yen5k <= (N - yen10k); yen5k++ {
			yen1k := N - (yen10k + yen5k)
			// fmt.Println(yen10k, yen5k, yen1k, price(yen10k, yen5k, yen1k))
			if price(yen10k, yen5k, yen1k) == Y {
				x, y, z = yen10k, yen5k, yen1k
			}
		}
	}
	return
}

func main() {
	input := bufio.NewReader(os.Stdin)

	var x, y int

	fmt.Fscan(input, &x, &y)

	x, y, z := solve(x, y)
	fmt.Println(x, y, z)
}
