package main

import "fmt"

func greeting(name string) string {
	return "Here we go - " + name
}

//func sum(a int, b int) int {
//	return a + b
//}

func sum(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(greeting("Karan"))
	fmt.Println(sum(1, 2))
}
