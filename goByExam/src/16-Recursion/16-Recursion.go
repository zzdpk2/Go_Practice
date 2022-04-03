package main

import "fmt"

// Write a fibonacci sequence
func main() {
	dummy := fibonacci(7)
	fmt.Println("dummy val: ", dummy)
}

func fibonacci(n int) int {
	// Base case
	if n == 0 {
		return 1
	}
	return n * fibonacci(n-1)
}
