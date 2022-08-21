package main

import "fmt"

func main() {
	fullName := "Aubdur Rob Anik"
	age := 15
	cgpa := 3.57
	university := "University Of Asia Pacific"

	//can't change this constant variable
	const COUNTRY = "Bangladesh"

	fmt.Printf("%v is a student\n", fullName)
	fmt.Printf("%v is %v years old\n", fullName, age)
	fmt.Printf("%v has cgpa is %v\n", fullName, cgpa)
	fmt.Printf("%v university name is %v\n", fullName, university)
	fmt.Printf("%v country is %v\n", fullName, COUNTRY)
}
