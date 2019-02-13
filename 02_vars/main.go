// Main types
// string
// bool
// int
// int8 - int64
// uint - uint64 & uintptr
// byte - alias for uint8
// rune - alias for int32
// float32 - float64
// complex64 complex 128

package main

import "fmt"

func main() {
	var name = "Dennis Onder"
	const isCool = true
	// Shorthand variable syntax
	size := 1.3
	fmt.Println(name, isCool, size)
	fmt.Printf("%T\n", name)
}
