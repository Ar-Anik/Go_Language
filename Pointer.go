package main

import "fmt"

func change(x int) {
	x = 100
}

func changev(x *int) {
	*x = 200
}

func main() {
	var p *int

	x := 10
	p = &x

	fmt.Printf("Memory Address of x : %v\n", &x)
	fmt.Printf("P : %v\n", p)
	fmt.Printf("Value of P : %v\n", *p)

	*p = *p + 2
	fmt.Printf("Value of X :: %v\n", x)

	change(x)
	fmt.Printf("After Change Function Call Value of X : %v\n", x)

	changev(&x)
	fmt.Printf("After Changev Function Call Value of X : %v\n", x)
}
