package main

import "fmt"

func change(x *int) {
	*x = 200
}

func main() {
	a := 10
	change(&a)

	fmt.Println(a)
}
