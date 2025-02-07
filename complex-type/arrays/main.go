package main

import "fmt"

func main()  {
	// Array declaration
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"

	// Array declaration
	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	fmt.Println("Number of colors:", len(colors))
	fmt.Println("Number of numbers:", len(numbers))
}