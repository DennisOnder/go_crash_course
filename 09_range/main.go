package main

import "fmt"

func main() {
	ids := []int{33, 76, 55, 12, 78, 3}
	// Loop through IDs
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}
	// Add IDs together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println(sum)
}
