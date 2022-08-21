package main

import (
	"fmt"
	"math"
)

func sum(num1, num2 float64) (result float64) {
	result = num1 + num2
	return
}

func sub(num1, num2 float64) (result float64) {
	result = math.Abs(num1 - num2)
	return
}

func mul(num1, num2 float64) (result float64) {
	result = num1 * num2
	return
}

func div(num1, num2 float64) (result float64) {
	result = num1 / num2
	return
}

func main() {
	i := true

	for i == true {
		var num1, num2 float64
		var option string

		fmt.Print("Enter First Number : ")
		fmt.Scan(&num1)

		fmt.Print("Enter Second Number : ")
		fmt.Scan(&num2)

		fmt.Print("Enter Option (+, -, *, /) : ")
		fmt.Scan(&option)

		switch option {
		case "+":
			fmt.Printf("Summation : %v\n", sum(num1, num2))
		case "-":
			fmt.Printf("Subtruction : %v\n", sub(num1, num2))
		case "*":
			fmt.Printf("Multiplication : %v\n", mul(num1, num2))
		case "/":
			fmt.Printf("Division : %v\n", div(num1, num2))
		default:
			fmt.Println("Invalid Input")
		}

	}
}
