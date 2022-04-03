package main

import "fmt"

func vals(a int, b int) (int, int) {
	return a, b
}

func main() {
	fmt.Println(vals(3, 7))
}
