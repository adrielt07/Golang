package main

import "os"

func main() {

     if len(os.Args) == 1 {
     	print("pass one Argument. 1 or 2\n")
     } else {
          switch {
	  case os.Args[1] == "1":
	       println("Foo")
     	  case os.Args[1] == "2":
	       println("Bar")
     	  }
     }
}