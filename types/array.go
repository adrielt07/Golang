package main

import "fmt"

// this is an array, Lenght is 10 and the values are int
var my_array = [10]int{12, 3, 4, 5, 6}

// Array of strings, but len is undeclared using variadic
var variadic_array = [...]string{"a", "b", "c"}

func main() {
	fmt.Printf("%v\n", my_array)
	fmt.Printf("%v\n", variadic_array)
}
