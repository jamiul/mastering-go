package main

import (
	"fmt"
	"sort"
)

func main()  {
	// Array declaration
	// var colors [3]string
	// colors[0] = "Red"
	// colors[1] = "Green"
	// colors[2] = "Blue"

	// Array declaration
	// var numbers = [5]int{1, 2, 3, 4, 5}
	// fmt.Println(numbers)

	// fmt.Println("Number of colors:", len(colors))
	// fmt.Println("Number of numbers:", len(numbers))

	// Managed ordered value in slice
	// This is a slice
	var colors = make([]string, 0, 3)
	colors = append(colors, "Red", "Green", "Blue")
	fmt.Println(colors)
	colors = append(colors, "Purple", "Yellow")
	fmt.Println(colors)

	colors = remove(colors, 1)
	fmt.Println(colors)

	sort.Strings(colors)
	fmt.Println(colors)
}

func remove(slice []string, i int) []string {
	return append(slice[:i], slice[i+1:]...)
}