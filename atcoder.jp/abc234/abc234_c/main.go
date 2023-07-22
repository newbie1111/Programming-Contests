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

func solve(k int) string {
	var (
		s   string
		res string
	)
	for shift := 0; 1<<shift <= k; shift++ {
		if k>>shift&1 == 1 {
			s += "2"
		} else {
			s += "0"
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}

	return res
}

func main() {
	var k int
	fmt.Scan(&k)
	ans := solve(k)
	fmt.Println(ans)
}
