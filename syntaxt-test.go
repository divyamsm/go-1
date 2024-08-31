package main

import (
	"fmt"
)

func main() {
	// Println prints with a newline
	fmt.Println("Hello, World!")

	// Printf formats according to a format specifier
	name := "Alice"
	age := 30
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	// Sprintf formats and returns a string
	message := fmt.Sprintf("Welcome, %s!", name)
	fmt.Println(message)
}
