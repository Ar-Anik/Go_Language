package main

import "fmt"

func main() {

	//Example-1
	var university, nowork string
	var cgpa float32

	fmt.Print("Enter University Name and CGPA : ")
	fmt.Scanf("%v %v", &university, &cgpa) //Enter : UAP 3.57

	fmt.Printf("University Name : %v\n", university)
	fmt.Printf("CGPA : %v\n", cgpa)

	fmt.Scanf("%v", &nowork) // it handle new line

	//Example-2
	var name string
	var age int
	var cg float32
	var status bool

	fmt.Print("Enter Name, age, cgpa and Married Status : ")
	fmt.Scanf("%s", &name)  // Sort Name
	fmt.Scanf("%d", &age)   // Integer Age
	fmt.Scanf("%g", &cg)    // Float cgpa
	fmt.Scanf("%t", status) // Married status--->True/False

	fmt.Printf("Name : %s\n", name)
	fmt.Printf("Age : %d\n", age)
	fmt.Printf("CGPA : %g\n", cgpa)
	fmt.Printf("Married Status : %t\n", status)
}
