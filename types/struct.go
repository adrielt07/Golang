package main

import "os"

// File that shows the usage of struct

type Salutation struct {
	name     string
	greeting string
}

func main() {
	var message Salutation

	if len(os.Args) > 1 {
		message = Salutation{os.Args[1], os.Args[2]}
	} else {
		message = Salutation{"Adriel", "Hello"}
	}

	println(message.greeting, message.name)
}
