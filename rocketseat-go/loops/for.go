package loops

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		// do something
	}

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, elem := range arr {
		fmt.Print(i, elem)
	}

	for i := range 10 {
		fmt.Print(i)
	}
}
