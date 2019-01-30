package main

import "fmt"
var (
    my_array = [4]int{1, 2, 3} // Array with fix size of 3
)

func main() {
     fmt.Println("Value of my_array", my_array, "\n") // Prints the value of my_array

     // Shows what happens to my_array if I created a copy of it using [:]
     var modified_my_array = my_array[:] // created "copy" of my_array
     modified_my_array[3] = 7
     fmt.Println("Modified copy of my array:", modified_my_array)
     fmt.Println("Original copy of my array:", my_array)

     fmt.Println("As shown in the example above. modified_my_array did not create 'copy' of the 'my_array'. On other words, they both reference the same array\n") 
}
