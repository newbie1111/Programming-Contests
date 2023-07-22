package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	input = bufio.NewScanner(os.Stdin)
	YES   = "YES"
	NO    = "NO"
)

func main() {
	var (
		n  int
		ab [][]int
	)
	input.Scan()
	n, _ = strconv.Atoi(input.Text())

	for i := 0; i < n; i++ {
		input.Scan()
		v := strings.Split(input.Text(), " ")
		av, _ := strconv.Atoi(v[0])
		bv, _ := strconv.Atoi(v[1])
		ab = append(ab, []int{av, bv})
	}

	type kv struct {
		Key   int
		Value float64
	}
	var ans []kv

	for i := 0; i < n; i++ {
		ans = append(ans, kv{i, float64(ab[i][0]) / float64(ab[i][0]+ab[i][1])})
	}

	sort.Slice(ans, func(i, j int) bool {
		return ans[i].Value > ans[j].Value
	})

	var output []int
	for _, v := range ans {
		output = append(output, v.Key+1)
	}

	for _, num := range output {
		fmt.Print(num, " ")
	}
	fmt.Println()
}
