package main

import "fmt"

func main() {
	ids := []int{11, 22, 33, 44, 55}

	// with index 'i'
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// without index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add
	sum := 0
	for _, id := range ids {
		sum += id
	}

	fmt.Println("Sum is: ", sum)

	// maps
	handle := map[string]string{"Ram": "@ramgo", "Shayam": "@shayamgo"}

	for k, v := range handle {
		fmt.Printf("%s: %s\n", k, v)
	}

	// Only keys
	for k := range handle {
		fmt.Println("Name: " + k)
	}

	// Only values
	for k := range handle {
		fmt.Println("Value: " + handle[k])
	}
}
