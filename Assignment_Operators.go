package main

import "fmt"

func main() {
	// Assignment Operators --> =, +=, -=, *=, /=, %=

	var x = 5
	fmt.Printf("X = %v\n", x)

	x += 3
	fmt.Printf("X = %v\n", x)

	x -= 5
	fmt.Printf("X = %v\n", x)

	x *= 2
	fmt.Printf("X = %v\n", x)

	x /= 3
	fmt.Printf("X = %v\n", x)

	x %= 4
	fmt.Printf("X = %v\n", x)
}
