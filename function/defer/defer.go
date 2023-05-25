package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}

// defer statement defers the execution of a function until the surrounding function returns.