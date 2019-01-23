package main

import "fmt"

func main() {
	// Define map
	emails := make(map[string]string)

	// Assign key value
	emails["Ram"] = "ram@go.com"
	emails["Shayam"] = "shayam@go.com"
	emails["Saurabh"] = "saurabh@go.com"

	fmt.Println("Emails: ", emails)
	fmt.Println(emails["Ram"])
	fmt.Println("Emails Size: ", len(emails))

	emails["Ravi"] = "ravi@go.com"
	fmt.Println("Emails: ", emails)
	fmt.Println("Emails Size: ", len(emails))

	// delete from map
	delete(emails, "Ravi")
	fmt.Println("Emails: ", emails)
	fmt.Println("Emails Size: ", len(emails))

	// Declare map and add kv
	handle := map[string]string{"Ram": "@ramgo", "Shayam": "@shayamgo"}

	fmt.Println("Twitter Handle: ", handle)
}
