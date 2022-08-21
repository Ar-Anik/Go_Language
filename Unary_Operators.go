package main

import "fmt"

func main() {
	// Unary Operators --> ++, --
	x := 3

	x++
	fmt.Printf("After Increment X = %v\n", x)

	x--
	fmt.Printf("After Decrement X = %v\n", x)

	/*  this give error
	Pre-Decrement : --x
	Pre-Increment : ++x
	*/
}
