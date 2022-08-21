package main

import "fmt"

func main() {

	//Example-1
	var num int

	fmt.Print("Enter Any Number : ")
	fmt.Scan(&num)

	if num > 0 {
		fmt.Println("Positive Number")
	} else if num < 0 {
		fmt.Println("Negative Number")
	} else {
		fmt.Println("Zero")
	}

	//Example-2
	var n int

	fmt.Print("Enter A Number : ")
	fmt.Scan(&n)

	if n%2 == 0 {
		fmt.Println("Even Number")
	} else {
		fmt.Println("Odd Number")
	}

	//Example-3
	var num1, num2, num3 int

	fmt.Print("Enter First Number : ")
	fmt.Scan(&num1)

	fmt.Print("Enter Second Number : ")
	fmt.Scan(&num2)

	fmt.Print("Enter Third Number : ")
	fmt.Scan(&num3)

	if num1 > num2 && num1 > num3 {
		fmt.Printf("%v is Large Number.\n", num1)
	} else if num2 > num1 && num2 > num3 {
		fmt.Printf("%v is Large Number.\n", num2)
	} else if num3 > num1 && num3 > num2 {
		fmt.Printf("%v is Large Number.\n", num3)
	} else {
		fmt.Printf("Number are Equal.\n")
	}
}
