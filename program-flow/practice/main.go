package main

import (
	"fmt"
)

type CartItem struct {
	Name     string
	Price    float64
	Quantity int
}

// calculateTotal() returns the total value of the shopping cart.
func calculateTotal(cart []CartItem) float64 {
	var total float64 = 0

	// for i := 0; i < len(cart); i++ {
	// 	itemTotal := cart[i].Price * float64(cart[i].Quantity)
	// 	total += itemTotal
	// }

	for _, item := range cart {
		total += (item.Price * float64(item.Quantity))
	}

	return total
}

func main() {
	var cart []CartItem
	var apples = CartItem{"apple", 2.99, 3}
	var oranges = CartItem{"orange", 1.50, 8}
	var bananas = CartItem{"banana", .49, 12}
	cart = append(cart, apples, oranges, bananas)
	fmt.Printf("Total: %.2f\n", calculateTotal(cart))
}
