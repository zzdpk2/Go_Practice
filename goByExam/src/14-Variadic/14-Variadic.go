package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("total: ", total)
}

func main() {
	sum(3, 4, 5)

	nums := []int{1, 2, 3}
	sum(nums...)
}
