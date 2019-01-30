package main

import "os"

func main() {

     switch {
     case os.Args[1] == "1":
     	  println("Foo")
     case os.Args[1] == "2":
     	  println("Bar")
     }
}