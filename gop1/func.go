package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		// return a / b
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s ", op)
	}
}

func div(a, b int) (q, r int) {
	return a / b, a % b
	// q = a / b
	// r = a % b
	// return
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling fucntion %s with args "+
		" (%d %d) ", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func swap1(a, b *int) {
	*a, *b = *b, *a
}

func swap2(a, b int) (int, int) {
	a, b = b, a
	return a, b
}

func main() {
	// fmt.Println(eval(3, 4, "*"))
	// fmt.Println(eval(3, 4, "x"))
	if result, err := eval(3, 5, "x"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	q, r := div(7, 8)
	fmt.Println(q, r)

	fmt.Println(apply(pow, 3, 4))

	fmt.Println(apply(
		func(a int, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 5))

	fmt.Println(sum(1, 2, 3, 4, 5))

	a, b := 3, 4
	// swap1(&a, &b)
	a, b = swap2(a, b)
	fmt.Println("a: ", a, " b: ", b)
}
