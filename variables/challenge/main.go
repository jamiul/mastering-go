/*
Create a Simple Calculator App
Task:
- Parse the parameters as float64 and return the sum of the two values.
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	value1 := "100.0"
	value2 := "-110.5"
	result := calculate(value1, value2)
	fmt.Printf("Summation: %.2f\n", result)
}

func calculate(value1 string, value2 string) float64  {
	// Convert the string to a float64
	num1, err := strconv.ParseFloat(value1, 64)
	if err != nil {
		return 0
	}

	num2, err := strconv.ParseFloat(value2, 64)
	if err != nil {
		return 0
	}

	return num1 + num2
}


