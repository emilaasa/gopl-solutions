package main

import (
	"fmt"
	"time"
)

// A first simple example of two functions running concurrently
// The purpose of the slow fib is to just take up a lot of time
func main() {
	go spinner(100 * time.Millisecond)
	fibN := fib(45)
	fmt.Printf("\rFibonacci(%d) = %d\n", 45, fibN)
}

func spinner(delay time.Duration) {
	for { // while true
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
