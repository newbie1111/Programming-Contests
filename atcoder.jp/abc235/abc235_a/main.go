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

func solve(a, b, c int) int {
	return 111 * (a + b + c)
}

func byte_to_int(r byte) int {
	return int(r - '0')
}

func main() {
	input.Scan()
	abc := input.Text()
	a, b, c := byte_to_int(abc[0]), byte_to_int(abc[1]), byte_to_int(abc[2])
	ans := solve(a, b, c)
	fmt.Println(ans)
}
