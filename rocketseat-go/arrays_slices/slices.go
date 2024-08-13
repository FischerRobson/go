package main

import (
	"fmt"
)

var products = []string{
	"Mouse",
	"Computer",
	"Keyboard",
	"Mousepad",
	"Monitor",
	"Laptop",
	"Acessories",
}

func main() {
	// ids := []int{1, 2, 3, 4, 5}

	// var selectedProducts = []string
	// selectedProducts := make([]string, 0, 5) // FIXED SIZE SLICE FOR BETTER MEMORY MANAGEMENT

	// for _, id := range ids {
	// 	prod := products[id]
	// 	fmt.Println(len(selectedProducts), cap(selectedProducts))
	// 	selectedProducts = append(selectedProducts, prod)
	// 	fmt.Println(len(selectedProducts), cap(selectedProducts))
	// }
	slice := []int{1, 2, 3, 4}
	foo(slice)
}

// func mai2() {
// 	arr := []int{1, 2, 3, 4, 5}
// 	slice := arr[:2]                 // COPY 1, 2 FROM arr
// 	fullSliceExpression := arr[:2:2] // COPY 1, 2 AND SET CAP TO 2
// 	fmt.Println(slice, cap(slice))
// 	fmt.Println(fullSliceExpression, cap(fullSliceExpression))
// }

func foo(slice []int) {
	_ = slice[3] // bound check
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice[2])
	fmt.Println(slice[3])
}
