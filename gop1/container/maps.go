package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "aa",
		"course":  "golang",
		"site":    "zebra",
		"quality": "not bad",
	}

	m2 := make(map[string]int) // m2 is empty map

	var m3 map[string]string // m3 is nil

	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map:")

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values:")
	courseName := m["course"]
	fmt.Println(courseName)
}
