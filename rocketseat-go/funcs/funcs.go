package funcs

import "fmt"

func defaultFunc() {
	fmt.Print("default func")
}

func sum(a int, b int) int {
	return a + b
}

// multi value return
func swap(a, b int) (int, int) {
	return b, a
}

// naked return - should not use
func divide(a, b int) (res int, rem int) {
	res = a / b
	rem = a % b
	return
}

// high order function
func highOrderSum(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

func sumN(nums ...int) int {
	var output = 0
	for _, i := range nums {
		output += i
	}
	return output
}
