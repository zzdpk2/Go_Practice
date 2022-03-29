package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map: ", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	fmt.Println("len: ", len(m))

	_, ok := m["k3"]
	fmt.Println(ok)

	if v2, ok := m["k3"]; ok {
		fmt.Println("v2:", v2)
	}

	n := map[int]string{1: "fuck huawei", 2: "hua wei is shithole"}
	fmt.Println(n)

}
