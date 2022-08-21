package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Print("Enter First Number : ")
	fmt.Scan(&num1)

	fmt.Print("Enter Second Number : ")
	fmt.Scan(&num2)

	r1 := num1 + num2
	fmt.Printf("Summation : %v\n", r1)

	r2 := num1 - num2
	fmt.Printf("Subtraction : %v\n", r2)

	r3 := num1 * num2
	fmt.Printf("Multiplication : %v\n", r3)

	r4 := num1 / num2
	fmt.Printf("Division : %v\n", r4)

	r5 := float32(num1) / float32(num2)
	fmt.Printf("Division : %v\n", r5)

	r6 := num1 % num2
	fmt.Printf("Modules : %v\n", r6)
}
