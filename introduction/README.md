# Introduction to Golang

### How to run?
```
go run <filename>
```
1. helloworld.go - prints "hello world!"
2. hellofriend.go - prints "hello world!" using println and prints "hello friend!" using fmt.Println
3. hi_first.go - prints "hi first before hello" using init. This shows that init function runs first before the main.
4. var_vs_init.go - prints "var goes first". This shows that variable gets declared first before init. Init can reassign the value of the variable as shown in this code
5. dont_change_me.go - prints "Edit this file and uncomment the 'message' (remove //) inside main() or init()". Go ahead and edit the file as stated. It shows that 'const' variables are constant (can't be reassigned with a different value).
6. slice.go - a simple example of slice
7. switch_case.go - an example use of switch and case statement. Takes one argument: 1 or 2
8. FizzBuzz.go - example usage of loop and if statements. Iterate up to 100. If num is a division of 3 print Fizz, if division of 5 print Buzz, lastly if it's a division of 15 print FizzBuzz.
9. basic_http_request - runs a simple http request in port 8000. Run it and open your browser. From your browser, enter 127.0.0.1:8000