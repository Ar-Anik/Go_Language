package main

import "fmt"

func main() {
	//Bitwise Operation --> &, |, ^, >>(Right Shift), <<(Left Shift)

	a := 5 // 5 = 101
	b := 7 // 7 = 111

	and := a & b //   = 101 = 5
	or := a | b  //   = 111 = 7
	xor := a ^ b //   = 010 = 2

	left_shift := a << 2  // 5 = 101 = 10100 = 20
	right_shift := b >> 2 // 7 = 111 = 001 = 1

	fmt.Printf("And Operation : %v\n", and)
	fmt.Printf("Or Operation : %v\n", or)
	fmt.Printf("Xor Operation : %v\n", xor)
	fmt.Printf("Left Shift : %v\n", left_shift)
	fmt.Printf("Right Shift : %v\n", right_shift)
}
