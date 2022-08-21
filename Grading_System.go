package main

import "fmt"

func main() {
	var n int

	fmt.Print("Enter a number between 0 to 100 : ")
	fmt.Scan(&n)

	if n >= 80 && n <= 100 {
		fmt.Println("A+")
	} else if n >= 70 && n <= 79 {
		fmt.Println("A")
	} else if n >= 60 && n <= 69 {
		fmt.Println("A-")
	} else if n >= 50 && n <= 59 {
		fmt.Println("B")
	} else if n >= 40 && n <= 49 {
		fmt.Println("C")
	} else if n >= 33 && n <= 39 {
		fmt.Println("D")
	} else {
		fmt.Println("Fail")
	}
}
