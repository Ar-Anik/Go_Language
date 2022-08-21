package main

import "fmt"

func main() {
	var decimalNumber int

	fmt.Print("Enter A Decimal Number : ")
	fmt.Scanf("%d", &decimalNumber)

	fmt.Printf("Binary Number : %b\n", decimalNumber)
	fmt.Printf("Octal Number : %o\n", decimalNumber)
	fmt.Printf("Hexa-Decimal Number : %x\n", decimalNumber)
}
