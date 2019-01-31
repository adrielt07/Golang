package main

import (
       "fmt"
)

func fizzbuzz(s int) {
     for i := 0; i <= s; i++ {
     	 if i % 15 == 0 {
	    fmt.Printf("Fizz")
	 } else if i % 5 == 0{
     	    fmt.Printf("Buzz")
	 } else if i % 3 == 0{
     	    fmt.Printf("FizzBuzz")
	 } else {
	    fmt.Printf("%d", i)
	 }

	 if i == s {
	    fmt.Printf("\n")
	 } else {
	    fmt.Print(", ")
	 }
     }
}

func main() {
     fizzbuzz(100)
}
