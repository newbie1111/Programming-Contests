package main

import (
	"fmt"
	"time"
)

func subtask() {
	for i := 0; i < 10; i++ {
		fmt.Println("subtask ", i)
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	go subtask()
	for i := 0; i < 10; i++ {
		fmt.Println("maintask ", i)
		time.Sleep(10 * time.Millisecond)
	}
}
