package main

import "fmt"

func main() {

	// var colors map[string]string

	colors := make(map[string]string)

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#00ff00",
	// 	"blue":  "#0000ff",
	// }

	// add element
	colors["red"] = "#ff0000"

	// remove element
	// delete(colors, "red")

	printMap(colors)

	fmt.Println(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color + ": " + hex)
	}
}
