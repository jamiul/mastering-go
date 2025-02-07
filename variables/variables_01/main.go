/*
Built-in  Integer Types
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

Alias  Integer Types
byte        alias for uint8
rune        alias for int32
uint        either 32 or 64 bits
int         same size as uint
unitptr     an unsigned integer large enough to store the uninterpreted bits of a pointer value

Floating-Point Types
float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

Complex Types
complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

Data Collection Types
array       a numbered sequence of elements of a single type with a fixed length
slice       an array-like type that can grow and shrink
map         a collection of key/value pairs
struct      a collection of fields of different types

Language Types
function    a sequence of statements
interface   a set of methods
channel     a communication mechanism that allows one goroutine to pass data to another goroutine

*/

// Output text to the console
package main

import (
	"fmt"
)

func main() {
	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."

	aNumber := 42
	fmt.Println(str1, str2, str3)
	stringLength, err := fmt.Println("The value is", aNumber)

	if err == nil {
		fmt.Println("The length of the string is", stringLength)
	}

	fmt.Printf("Value of number: %v\n", aNumber)

	fmt.Printf("Data type: %T\n", aNumber)
}