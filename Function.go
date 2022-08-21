package main

import "fmt"

func dispaly() {
	fmt.Println("My Country Name is Bangladesh")
}

func square(num float32) {
	fmt.Printf("Square : %v\n", num*num)
}

func sum(num1 float32, num2 float32) float32 {
	result := num1 + num2
	return result
}

func sub(num1 float32, num2 float32) float32 {
	var result float32

	if num1 >= num2 {
		result = num1 - num2
	} else {
		result = num2 - num1
	}

	return result
}

func multi(num1 float32, num2 float32) (result float32) {
	result = num1 * num2
	return
}

func div(num1, num2 float32) (result float32) {
	result = num1 / num2
	return
}

func main() {
	dispaly()
	square(5)

	fmt.Printf("Summation : %v\n", sum(10, 20))
	fmt.Printf("Subtruction : %v\n", sub(10, 20))
	fmt.Printf("Multiplication : %v\n", multi(5, 8))
	fmt.Printf("Division : %v\n", div(77, 7))
}
