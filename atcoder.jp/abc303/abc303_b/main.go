package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	type pair struct {
		a int
		b int
	}

	ans := 0
	person := make(map[pair]bool)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)

	input.Scan()
	n, _ := strconv.Atoi(input.Text())
	input.Scan()
	m, _ := strconv.Atoi(input.Text())

	for j := 0; j < m; j++ {
		var data []int
		for i := 0; i < n; i++ {
			input.Scan()
			a, _ := strconv.Atoi(input.Text())
			data = append(data, a)
		}

		for i := 0; i < n-1; i++ {
			if data[i] < data[i+1] {
				person[pair{data[i], data[i+1]}] = true
			} else {
				person[pair{data[i+1], data[i]}] = true
			}
		}
	}

	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			if !person[pair{i, j}] {
				ans++
			}
		}
	}

	fmt.Println(ans)

}
