package main

import "fmt"

func main() {
	x := 20
	y := 20
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else if x == y {
		fmt.Printf("%d is equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is greater than %d", x, y)
	}
	// Switch
	color := "green"
	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Color is not neither red or blue")
	}
}
