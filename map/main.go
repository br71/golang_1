package main

import "fmt"

// func main() {
// 	// [string = type of keys]string = type of valyes
// 	// var colors map[string]string

// 	colors := make(map[string]string)

// 	// colors := map[string]string{
// 	// 	"red":   "#ff0000",
// 	// 	"green": "#4bf745",
// 	// }

// 	// Add values

// 	colors["white"] = "#ffffff"
// 	fmt.Println(colors)

// 	delete(colors, "white")
// 	fmt.Println(colors)

// }

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	// fmt.Println(colors)
	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
