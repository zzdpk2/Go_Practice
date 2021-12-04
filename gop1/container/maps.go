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

	causeName := m["cause"]
	fmt.Println("cause val", causeName)

	if causeName1, ok := m["cause"]; ok {
		fmt.Println("cause2 val: ", causeName1)
	} else {
		fmt.Println("key does not exist")
	}

	// Delete values
	fmt.Println("Deleting values:")
	name, ok := m["name"]
	fmt.Println(name, ok)

	// After deleting
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println("name:", name, ok)

	// What elements cannot be keys? The elements cannot be evaluated such as slice, map, and functions.
	// Struct without those data structures can be the key as well.
}
