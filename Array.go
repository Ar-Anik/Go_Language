package main

import "fmt"

func main() {
	//Example-1
	var country = [5]string{"Bangladesh", "Nepal", "Dubai", "China", "Japan"}
	for i := 0; i < len(country); i++ {
		fmt.Println(country[i])
	}

	//Example-2
	var num [10]int
	sum := 0
	fmt.Print("Enter 10 Number : ")
	for i := 0; i < 10; i++ {
		fmt.Scan(&num[i])
		sum += num[i]
	}

	fmt.Printf("Summation : %v\n", sum)
}
