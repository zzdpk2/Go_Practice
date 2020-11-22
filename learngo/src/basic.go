package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa = 3
	ss = "kkk"
	bb = true
)

func varZeroVal() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func varInitVal() {
	a, b, c, d := "a", 1, 3, true
	b = 5
	fmt.Println(a, b, c, d)
}

func euler() {
	// c := 3 + 4i
	// fmt.Print(cmplx.Abs(c))
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
}

func triangle() {
	a, b := 3, 4
	c := int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	c := int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	//const(
	//	cpp = 0
	//	java = 1
	//	python = 2
	//	golang = 3
	//)
	const (
		cpp = iota
		_
		java
		python
		golang
	)
	fmt.Println(cpp, java, python, golang)

	// b, kb, mb, gb, tb, pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)

}

func main() {
	fmt.Println("Hello World!")
	varZeroVal()
	varInitVal()
	euler()
	triangle()
	consts()
	enums()
	fmt.Println(aa, ss, bb)
}
