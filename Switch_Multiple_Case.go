package main

import "fmt"

func main() {
	var cls int

	fmt.Print("What class are you read : ")
	fmt.Scan(&cls)

	switch cls {
	case 1, 2, 3, 4, 5:
		fmt.Println("Primary")
	case 6, 7, 8, 9, 10:
		fmt.Println("Secondary")
	case 11, 12:
		fmt.Println("Higher-Secondary")
	default:
		fmt.Println("Invalid Input")
	}
}
