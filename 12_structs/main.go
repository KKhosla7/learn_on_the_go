package main

import (
	"fmt"
	"strconv"
)

// Define Employee struct
type Employee struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

// say hello (Value Receiver)
func (e Employee) sayHello() string {
	return "Hello, I'm " + e.firstName + ", " + strconv.Itoa(e.age) + " years old."
}

// birthday today (Pointer Receiver)
func (e *Employee) birthdayToday() {
	e.age++
}

// married today (Pointer Receiver)
func (e *Employee) marriedToday(lastName string) {
	if e.gender == "M" {
		return
	} else {
		e.lastName = lastName
	}
}

func main() {
	employee1 := Employee{firstName: "Ram", lastName: "Sharma", city: "New Delhi", gender: "M", age: 23}
	employee2 := Employee{"Shruti", "Verma", "Uttar Pradesh", "F", 24}
	fmt.Println(employee1, employee2)
	fmt.Println(employee2.firstName, employee1.lastName)

	employee1.birthdayToday()
	employee1.birthdayToday()
	fmt.Println(employee1.sayHello())

	employee1.marriedToday("Kaushal")
	employee2.marriedToday("Rana")
	fmt.Println(employee1)
	fmt.Println(employee2)
}
