package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "RED",
		"green": "GREEN",
		"white": "WHITE",
	}

	printMap(colors)

	// var colors map[string]string

	// colors := make(map[string]string)

	// colors["white"] = "WHITE"
	// delete(colors, "whitw") //does nothing

	// fmt.Println(colors)

	// delete(colors, "white")

	// fmt.Println(colors)
}

func printMap(c map[string]string) {
	for key, val := range c {
		fmt.Println(key, val)
	}
}
