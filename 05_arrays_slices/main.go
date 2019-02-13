package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string
	// Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"
	fmt.Println(fruitArr)
	// Declare and assign
	newArr := [2]string{"New", "Array"}
	fmt.Println(newArr)
	// Slices
	fruitSlice := []string{"Orange", "Grapes", "Bananas"}
	fmt.Println(fruitSlice)
}
