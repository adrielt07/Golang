package main

import "fmt"

// Pointer points to an address
// Pointers are specific to its type. For example, int pointer can't point to a string

func main() {
	// Declared an int
	var num int = 45

	// Pointer to an int. Assign the pointer to the memory address by using &
	var pointer *int = &num

	// dereference pointer
	fmt.Printf("Memory address of num: %v \nValue of num: %v\n", pointer, *pointer)
}
