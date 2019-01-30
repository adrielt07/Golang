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
