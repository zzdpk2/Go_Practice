package main

import "fmt"

/*
#include <stdio.h>
void Print(){
	printf("%d\n", 666);
}
*/
import "C"

func main() {
	fmt.Print("Hi!")
	fmt.Print("New Program")

	var dummyHead int = 123
	sum := dummyHead + 6
	fmt.Print("Sum:", sum)
	fmt.Print("dummyHead:", dummyHead)
	fmt.Print("Program successfully compiled!")

	// Call C functions in Go
	C.Print()

}
