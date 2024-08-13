package main

import "fmt"

func main() {
	var m map[string]string
	m["name"] = "robson" // generate a panic as map is nil

	// map initialization
	m2 := make(map[string]string)
	m2["name"] = "robson"

	// map initialization
	m3 := map[string]string{
		"name": "robson",
	}

	m4 := map[string][]int{
		"robson": []int{1, 2, 3}, // type can be omitted
		"john":   {2, 3, 4},
	}

	for key, value := range m4 {
		fmt.Println(key, value)
	}

	fmt.Println(m2)
	fmt.Println(m3)
	fmt.Println(m4)
}
