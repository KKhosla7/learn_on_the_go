package main

import "fmt"

func main() {
	// Use & to print the memory address
	// Use * to read the value from memory address
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Println(*b, *&a)

	// change val using pointer
	*b = 10
	fmt.Println(a)
}
