package main

import "fmt"

func main() {
	x := 42
	p := &x // Get the pointer to x

	fmt.Println("x:", x)   // 42
	fmt.Println("p:", p)   // Memory address of x
	fmt.Println("*p:", *p) // 42 (value at the memory address p)

	*p = 100                            // Modify x through the pointer
	fmt.Println("x after *p = 100:", x) // 100

	// -----------------------------------------------

	y := 10
	take(&y)
	fmt.Println(y)
}

func take(number *int) {
	*number = 100
}
