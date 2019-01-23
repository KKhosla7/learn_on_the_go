package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string

	// Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// Declare and Assign
	vegArr := [2]string{"Carrot", "Reddish"}

	// Slices
	playerSlice := []string{"Kholi", "Dhoni", "Jadeja", "Sharma"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])

	fmt.Println(vegArr)

	fmt.Println(playerSlice)
	fmt.Println(len(playerSlice))
	fmt.Println(playerSlice[1:3])
}
