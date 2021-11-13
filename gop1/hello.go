package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

const filename = "abc.txt"

var (
	aa, ss, bb = 3, "kkk", true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

func euler() {
	fmt.Printf("%.3f\n",
		//cmplx.Pow(math.E, 1i*math.Pi) + 1,
		cmplx.Exp(1i*math.Pi)+1,
	)
}

func triangle() {
	var a, b int = 3, 4
	var c int = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func consts() {
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enum() {
	const (
		cpp = iota
		_
		python = 2
		golang = 3
	)
	//const (
	//	cpp    = 0
	//	java   = 1
	//	python = 2
	//	golang = 3
	//)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	//fmt.Println(runtime.GOARCH)
	//fmt.Println("Hello World!")
	//variableZeroValue()
	//variableInitialValue()
	//variableTypeDeduction()
	//variableShorter()
	euler()
	triangle()
	consts()
	enum()
}
