package main

import "fmt"

func zeroVal(val int) {
	val = 0
}

func zeroPointer(p *int) {
	*p = 0
}

func main() {
	i := 1
	fmt.Println("Init value: ", i)

	zeroVal(i)
	fmt.Println("Value1: ", i)

	zeroPointer(&i)
	fmt.Println("Value2: ", i)

	fmt.Println("Address: ", &i)

}
