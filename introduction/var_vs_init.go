package main

var (
    message string = "Hello World!"
)

func main() {
     println(message)
}

// vars will run first before the init.
// Therefore, value of message will get reassigned to "var goes first"
func init() {
     println("Variables declared in var block will get instantiated first")
     message = "var goes first"
}
