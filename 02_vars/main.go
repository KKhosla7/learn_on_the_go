package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	// var name = "Karan"

	var age = 37
	var isCool = true
	// var size = 2.5 // By default float64

	// Shorthand
	name := "Karan"
	size := 1.3

	a, b := 1, 2

	const immutable = "YES"

	fmt.Println(name, age, isCool, size, a, b, immutable)
	fmt.Printf("%T\n", isCool)
}
