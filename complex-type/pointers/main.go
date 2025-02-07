/*
This code demonstrates basic pointer operations in Go, including:

Pointer declaration (*type)
Address-of operator (&)
Dereferencing operator (*)
Pointer/Value relationship
*/
package main

import (
	"fmt"
)

func main()  {
	anInt := 42
	//p is declared as a pointer to int (*int) and assigned the address of anInt using &
	var p *int = &anInt

	if p == nil {
		fmt.Println("p is nill")
	} else {
		// *p accesses the value stored at the memory address
		fmt.Println("Value of p:", *p)
	}

	value1 := 42.23
	pointer1 := &value1
	*pointer1 = *pointer1

	fmt.Println("Pointer:", *pointer1)
	fmt.Println("Original Value:", value1)
}