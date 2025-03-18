package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	value1 := "100"
	value2 := "3"
	operation := "/"
	result := calculate(value1, value2, operation)
	fmt.Printf("Result: %v\n", result)
}

func calculate(input1 string, input2 string, operation string) float64 {
	num1 := convertInputToValue(input1)
	num2 := convertInputToValue(input2)

	var result float64

	// if operation == "+" {
	// 	result = addValues(num1, num2)
	// }

	// if operation == "-" {
	// 	result = subtractValues(num1, num2)
	// }

	// if operation == "*" {
	// 	result = multiplyValues(num1, num2)
	// }

	// if operation == "/" {
	// 	result = divideValues(num1, num2)
	// }

	// Another apporach
	switch operation {
	case "+":
		result = addValues(num1, num1)
	case "-":
		result = subtractValues(num1, num2)
	case "*":
		result = multiplyValues(num1, num2)
	case "/":
		result = divideValues(num1, num2)
	default:
		panic("Invalid operation")
	}
	return result
}

// func convertInputToValue(input string) float64 {
// 	value, err := strconv.ParseFloat(input, 64)

// 	if err != nil {
// 		fmt.Println("Error converting", input, "to float", err)
// 		return 0
// 	}
// 	return value
// }

// Alternative apporach
func convertInputToValue(input string) float64 {
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		message := fmt.Sprintf("%v must be a number", input)
		panic(message)
	}
	return value
}

func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	if value2 != 0 {
		return value1 / value2
	}

	fmt.Println("Error: Division by zero")
	return 0
}