package main

import "fmt"

// change_value - changes the value of an int
// num - pointer to an int
func change_value(num *int) {
	*num = 123456
}

// An example of how the value of a variable changes when
// its memory address is passed to a pointer
func main() {
	var number1s = 11111

	fmt.Printf("Value of number1s before: %v\n", number1s)
	change_value(&number1s)
	fmt.Printf("Value of number1s After: %v\n", number1s)
}
