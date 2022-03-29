package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	fmt.Println("sum: ", sum)

	for i, v1 := range nums {
		if i == 2 {
			fmt.Println("v1：", v1)
		}
	}

	n := map[int]string{1: "a", 2: "b", 3: "c", 4: "d"}
	for k, v := range n {
		fmt.Printf("key: %d, val: %s\n", k, v)
		if k == 2 {
			break
		}
	}

	for k, v := range "niu pi" {
		fmt.Println("key: ", k)
		fmt.Println("val: ", v)
	}
}
