// Using math operators
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	num1, num2, num3 := 10, 20, 30
// 	intSum := num1 + num2 + num3
// 	fmt.Println("Integer sum:", intSum)

// 	num4, num5, num6 := 10.5, 20.5, 30.5
// 	floatSum := num4 + num5 + num6
// 	fmt.Println("Float sum:", floatSum)

// 	total := float64(num1) + num4
// 	fmt.Println("Total:", total)
// }

// Using math package
package main

import (
	"fmt"
	"math"
)

func main() {
	fNum1, fNum2 := 10.56, 20.55
	floatSum := fNum1 + fNum2
	fmt.Println("Float sum:", floatSum)

	sum := math.Round(floatSum * 100) / 100
	fmt.Printf("Rounded sum: %v \n", sum)

	fmt.Println("The value of Pi:", math.Pi)

	circleRadius := 15.5
	circumference := 2 * math.Pi * circleRadius
	fmt.Printf("Circumference: %.2f \n", circumference)
}