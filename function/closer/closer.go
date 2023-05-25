package main

import "fmt"

// limit scope

var x int

func main() {
	fmt.Println(x)
	x++
	fmt.Println(x)
	fmt.Print(x)

	{
		y:= 42
		fmt.Println(y)
	}

	foo()
}

func foo() {
	fmt.Println("hello")
	x++
	// y=12 can't do this because y is not in scope
}	