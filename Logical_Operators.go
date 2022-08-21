package main

import "fmt"

func main() {
	// Logical Operators --> And = &&, Or = ||, Not = !

	x := 5
	y := 2

	var answer = x > 3 && y > 3
	fmt.Printf("And Operation : %v\n", answer)

	var result = x > 3 || y > 3
	fmt.Printf("Or Operation : %v\n", result)

	var response = !(x < 3)
	fmt.Printf("Not Operation : %v\n", response)
}
