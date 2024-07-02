package main

import "fmt"

func main() {

	// colors := make(map[string]string)
	// ==
	// var colors map[string]string

	colors := map[string]string{
		"red":   "#ff0000",
		"black": "#000000",
		"white": "#ffffff",
	}

	colors["green"] = "#00ff00"

	delete(colors, "red")

	printMap(colors)
}

func printMap(ma map[string]string) {
	for color, hex := range ma {
		fmt.Println(color, hex)
	}
}
