package main

import "fmt"

func main() {
	// Define a map
	emails := make(map[string]string)
	// Assign values to map
	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"
	emails["Dennis"] = "dennis@gmail.com"
	emails["Delete"] = "delete@gmail.com"
	fmt.Println(emails)
	// Delete
	delete(emails, "Delete")
	fmt.Println(emails)
	// Declare map and add values
	emails2 := map[string]string{"Brad": "brad@gmail.com", "Ralph": "ralph@gmail.com"}
	fmt.Println(emails2)
}
