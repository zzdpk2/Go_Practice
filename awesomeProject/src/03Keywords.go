package main

import (
	"fmt"
	"math"
	"reflect"
)

func main01() {
	var a11 int = 10
	var b11 = 20
	// type induction
	c11 := 30

	k1, k2, k3 := 1, true, "Shadiao"
	fmt.Println(a11)
	fmt.Println(b11)
	fmt.Println(c11)

	fmt.Println(k1)
	fmt.Println(k2)
	fmt.Println(k3)

	// find out the tpe
	fmt.Printf("%T\n", k1)
	fmt.Printf("%T\n", k2)
	fmt.Printf("%T\n", k3)

	// Switching the values
	a1, a2 := 123, 666
	fmt.Println(a1, a2)
	a1, a2 = a2, a1
	fmt.Println(a1, a2)

	// const PI float64 = 3.1415926
	// fmt.Println(PI)
	fmt.Print(math.Pi)

	// ITOA sets (enum)
	// const (
	//	 a = iota
	//	 b = iota
	//	 c, d, e = iota, iota, iota
	// )
	const (
		a = true
		b = 123
		c = "hfsdfsd"
		d = 1.23
		e = 3e6
	)
	fmt.Println(a, b, c, d, e)

	func(a int, b int) {
		fmt.Println(a + b)
	}(10, 20)

	// Pre-declare
	var f1 func(int, int)
	f1 = func(a int, b int) {
		fmt.Println(a + b)
	}

	f1(10, 15)
	fmt.Printf("%T\n", f1)

	// Without Pre-declare
	f2 := func(a int, b int) {
		fmt.Println(a + b)
	}
	fmt.Printf("%T\n", f2)
}

// self define function type
type FUNCTYPE func(int, bool, string)

func demo(a int, b bool, c string) {
	fmt.Println("new function was created!")
}

// function recall
func demo2(f FUNCTYPE) {
	f(1, true, "Good")
}

func main02() {
	// type
	// 1. create alias for existed type
	main01()
	type i8 = int8
	//var a i8 = 123
	//var b i8 = 2
	var b i8 = 2
	var a int8 = 1
	fmt.Println(a + b)

	var f FUNCTYPE
	f = demo
	f(1, false, "good")

	var f2 FUNCTYPE
	f2 = demo
	demo2(f2)
}

func main03() {
	// reflection and interface
	// Empty type of interface that can accept all types of params
	var a interface{}
	a = 1
	a = 12.3
	a = true
	a = 1
	// Reflection operation to get type of a interface
	t := reflect.TypeOf(a)
	fmt.Println(t)
	// Reflection operation to get value of a interface
	// The type os v is not int instead it is reflect.value
	a1, ok := a.(int)
	fmt.Println(a1)
	fmt.Println(ok)
	fmt.Printf("a1 type: %T\n", a1)
	fmt.Printf("ok type: %T\n", ok)

	v := reflect.ValueOf(a)
	fmt.Println(v)
	fmt.Println(v.Kind())
	fmt.Println(v.Int())
	a = "Niu bi!"
	fmt.Println(a)
}

func main() {

}
