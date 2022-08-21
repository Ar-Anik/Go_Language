package main

import "fmt"

type Student struct {
	id       int
	semester string
	section  string
	cgpa     float32
	age      int
}

func display(str Student) {
	fmt.Printf("ID : %v\n", str.id)
	fmt.Printf("Semester : %v\n", str.semester)
	fmt.Printf("Section : %v\n", str.section)
	fmt.Printf("CGPA : %v\n", str.cgpa)
	fmt.Printf("Age : %v\n\n", str.age)
}

func (str *Student) ageIncrease(value int) {
	str.age = str.age + value
}

func main() {

	anik := Student{173, "4-2", "B", 3.57, 14}

	oni := Student{179, "4-2", "B", 3.52, 18}

	joy := Student{193, "4-2", "B", 3.73, 19}

	fmt.Printf("Anik ID : %v\n", anik.id)
	fmt.Printf("Anik Semester : %v\n", anik.semester)
	fmt.Printf("Anik Section : %v\n", anik.section)
	fmt.Printf("Anik CGPA : %v\n", anik.cgpa)
	fmt.Printf("Anik Age : %v\n\n", anik.age)

	display(oni)
	display(joy)

	anik.ageIncrease(5)
	display(anik)
}
