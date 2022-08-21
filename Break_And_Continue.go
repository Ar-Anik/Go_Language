package main

import "fmt"

func main() {

	i := 0

	for i <= 10 {
		i++

		if i%2 == 0 {
			continue
		}

		fmt.Printf("%v ", i)

		if i == 7 {
			break
		}
	}
}
