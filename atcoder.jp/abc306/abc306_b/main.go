package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")

	// 入力文字列をスペースで分割して配列に格納
	strNumbers := strings.Split(input, " ")

	var numbers []int
	for _, strNum := range strNumbers {
		num, _ := strconv.Atoi(strNum)
		numbers = append(numbers, num)
	}

	var ans uint64

	for i, v := range numbers {
		ans += uint64(v) * uint64(math.Pow(2, float64(i)))
	}

	fmt.Println(uint64(ans))
}
