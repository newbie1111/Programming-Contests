package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var input = bufio.NewScanner(os.Stdin)

// func dump(variable ...any) {
// 	for _, v := range variable {
// 		fmt.Printf("%v : %v\n", reflect.TypeOf(v), v)
// 	}
// }

func manhattan_distance_2d(sx, sy, to_x, to_y int) int {
	return int(float64(math.Abs(float64(sx-to_x))) + math.Abs(float64(sy-to_y)))
}

func solve(n int, t, x, y []int) string {
	now_t, now_x, now_y := 0, 0, 0
	Yes, No := "Yes", "No"

	for i := 0; i < n; i++ {
		// dump(t[i], now_t, t[i]-now_t, manhattan_distance_2d(now_x, now_y, x[i], y[i]))
		t_dist := t[i] - now_t
		xy_dist := manhattan_distance_2d(now_x, now_y, x[i], y[i])
		if xy_dist <= t_dist && int(math.Abs(float64(t_dist-xy_dist)))%2 == 0 {
			now_t = t[i]
			now_x, now_y = x[i], y[i]
		} else {
			return No
		}
	}

	return Yes
}

func main() {

	var (
		n int
		t []int
		x []int
		y []int
	)

	input.Scan()
	n, _ = strconv.Atoi(input.Text())
	for i := 0; i < n; i++ {
		input.Scan()
		list := strings.Split(input.Text(), " ")

		next_t, _ := strconv.Atoi(list[0])
		next_x, _ := strconv.Atoi(list[1])
		next_y, _ := strconv.Atoi(list[2])

		t = append(t, next_t)
		x = append(x, next_x)
		y = append(y, next_y)

	}
	// dump(n, t, x, y)

	fmt.Println(solve(n, t, x, y))
}
