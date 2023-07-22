package main

import (
	"bufio"
	"bytes"
	"container/list"
	"fmt"
	"os"
	"reflect"
)

var (
	input = bufio.NewScanner(os.Stdin)
	YES   = "YES"
	NO    = "NO"
)

// stack

type Stack struct {
	s *list.List
}

func stack() *Stack {
	return &Stack{s: list.New()}
}

func (s *Stack) push(val interface{}) {
	s.s.PushBack(val)
}

func (s *Stack) pop() interface{} {
	if tail := s.s.Back(); tail == nil {
		return tail
	} else {
		return s.s.Remove(tail)
	}
}

func reversed(iter []byte) {
	for i, j := 0, len(iter)-1; i < j; i, j = i+1, j-1 {
		iter[i], iter[j] = iter[j], iter[i]
	}
}

func dump(variable ...interface{}) {
	for _, v := range variable {
		fmt.Printf("%v : %v\n", reflect.TypeOf(v), v)
	}
}

func solve(s string) string {
	var (
		words = [][]byte{
			[]byte("dream"),
			[]byte("dreamer"),
			[]byte("erase"),
			[]byte("eraser"),
		}
		length = 0
		target = []byte(s)
	)
	for _, v := range words {
		reversed(v)
	}
	reversed(target)

	fmt.Println(target)

	for length < len(target) {
		for _, word := range words {
			fmt.Println(word, string(word))
			fmt.Println(length, target[length:], target[:length])
			if bytes.HasPrefix(target[length:], word) {
				length += len(word)
				fmt.Println(length, target[length:], target[:length], "\n")
			} else {
				return NO
			}
		}
	}

	return YES
}

func main() {
	input.Scan()
	s := input.Text()
	fmt.Println(solve(s))
}
