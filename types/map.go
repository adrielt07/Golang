package main

// Map is a type that has a key and a value

import "fmt"

func main() {

	/*
		var dictionary map[string]string
		dictionary = make(map[string]string)

		dictionary["key"] = "value"
		dictionary["number"] = "123"
		dictionary["alpha"] = "abc"
	*/

	// Shorthand declaration of a map
	dictionary := map[string]string{
		"Car":    "Toyota",
		"Shoe":   "Nike",
		"Sport":  "Basketball",
		"Number": "28",
	}

	dictionary["Car"] = "Nissan" // Update the value of key: Car
	delete(dictionary, "Shoe")   // Delete Shoe key and its value

	for k, v := range dictionary {
		fmt.Println("Key is:", k)
		fmt.Println("Value is:", v)
	}
}
