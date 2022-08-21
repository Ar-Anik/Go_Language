package main

import "fmt"

func main() {

	//Example-1
	fmt.Print("Odd Number : ")
	for i := 1; i <= 100; i += 2 {
		fmt.Printf("%v ", i)
	}
	fmt.Println()

	//Example-2
	fmt.Print("Even Number : ")
	i := 2
	for i <= 100 {
		fmt.Printf("%v ", i)
		i += 2
	}
}
