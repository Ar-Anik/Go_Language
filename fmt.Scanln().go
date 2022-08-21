package main

import "fmt"

func main() {
	var name, university string

	fmt.Print("Enter name and university : ")
	fmt.Scanln(&name, &university) //Anik UAP

	fmt.Printf("Name : %s\n", name)
	fmt.Printf("University : %s", university)
}
