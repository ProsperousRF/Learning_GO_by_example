package main

import (
	"fmt"
	"maps"
)

func main() {

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v3 := m["something doesn't existed"] // here will be 0
	fmt.Println("v3:", v3)

	fmt.Println("map length:", len(m))

	fmt.Println("Before delete:", m)
	delete(m, "k2")
	fmt.Println("After delete:", m)

	clear(m)
	fmt.Println("After clear:", m)

	theValue, prs := m["missing key"]
	fmt.Printf("The value: %d, prs: %t\n", theValue, prs)

	n := map[string]uint{"foo": 1, "bar": 2}
	fmt.Println("map: ", n)

	n2 := map[string]uint{"foo": 1, "bar": 2}
	fmt.Println(maps.Equal(n, n2))

}
