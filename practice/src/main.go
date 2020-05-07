package src

import "fmt"

func demo(a int, b int) int {
	return a + b
}

func main() {
	a := 3
	b := 4
	value := demo(a, b)
	fmt.Println(value)
}
