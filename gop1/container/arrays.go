package main

import "fmt"

func printArray(arr [5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArray2(arr *[5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
	arr[0] = 100
}

func printArray3(arr []int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
	arr[0] = 100
}

func main() {
	// var arr1 [5]int
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	// for i := 0; i < len(arr3); i++ {
	// 	fmt.Println(arr3[3])
	// }

	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	// Find out max value and its index
	numbers := [...]int{3, 7, 1, 9, 10, 78, 34}
	maxi, maxValue := -1, -1
	for i, v := range numbers {
		if v > maxValue {
			maxi, maxValue = i, v
		}
	}
	fmt.Println("maxi is : ", maxi)
	fmt.Println("maxValue is :", maxValue)

	printArray(arr3)
	printArray2(&arr3)
	fmt.Println(arr3)

	// Convert to slice
	printArray3(arr3[:])
}
