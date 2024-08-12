package main

import "fmt"

func main() {
	doDefer()
}

func doDefer() {
	x := 10
	defer func() {
		fmt.Println(x)
	}()

	x = 50
	fmt.Println(x)
}
