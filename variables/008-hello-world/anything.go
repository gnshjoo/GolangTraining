package main

import "fmt"

func main() {
	fmt.Println("Hello everyone, this is the most and learning the Go programming language.")
	foo()
	fmt.Println("something more")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}

func foo() {
	fmt.Println("i'm in foo")
}

func bar() {
	fmt.Println("and then we exited")
}

// control flow:
// 1. sequence
// 2. loop : iterative
// 3. conditional
