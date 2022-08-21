package main

import "fmt"

func main() {
	//Area of a Square
	var side float32

	fmt.Print("Enter a side of Square : ")
	fmt.Scan(&side)

	area := side * side
	fmt.Printf("Area of a Square : %v\n", area)

	//Area of a Triangle
	var base, height float32

	fmt.Print("Enter a Base : ")
	fmt.Scan(&base)

	fmt.Print("Enter a Height : ")
	fmt.Scan(&height)

	area1 := 0.5 * base * height
	fmt.Printf("Area of a Triangel : %v\n", area1)

	//Area of a Rectangle
	var length, width float32

	fmt.Print("Enter Length : ")
	fmt.Scan(&length)

	fmt.Print("Enter width : ")
	fmt.Scan(&width)

	area2 := length * width
	fmt.Printf("Area of a Rectangel : %v", area2)
}
