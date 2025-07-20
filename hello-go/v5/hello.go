package main

import "fmt"

const englishHelloPrefix = "Hello, "

// display Hello, World if an empty string is passed
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello(""))
}