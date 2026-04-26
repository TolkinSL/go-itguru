package main

import "fmt"

func main() {
	var a1 int8 = 3
	var b1 uint8 = 3

	fmt.Printf("^a, %d | %08b\n", ^a1, ^a1)
	fmt.Printf("^b, %d | %08b\n", ^b1, ^b1)

	var a2 int8 = 4
	fmt.Printf("a << 1, %d | %08b\n", a2<<1, a2<<1)

	var a3 int8 = 4
	fmt.Printf("a >> 1, %d | %08b\n", a3>>1, a3>>1)

	//четное не четное
	var a4 int8 = 4
	fmt.Printf("a & 1, a = %08b | %08b\n", a4, a4 & 1)

	a4 = 7
	fmt.Printf("a & 1, a = %08b | %08b\n", a4, a4 & 1)

	var a5 int8 = 1
	fmt.Printf("a << 7, %d | %08b\n", a5<<7, a5<<7)

	a5 = 1
	fmt.Printf("a << 6, %d | %08b\n", a5<<6, a5<<6)

}
