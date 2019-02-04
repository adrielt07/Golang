package main

// Importing packages located in example_package
import (
	"./examplepkg"
)

func main() {
	// Referencing MyAdd function from example_package
	var result int

	// Referencing Brands type from example_package
	var brand = examplepkg.Brands{"Shoe", "Nike"}

	result = examplepkg.MyAdd(1, 2, 3, 4, 5)

	println(result)
	println(brand.Name, brand.Value)
}
