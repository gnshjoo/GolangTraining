package main

import "fmt"

var y = 43

// DECLARE the variable with the IDENTIFIER "z"
// and that the variable with the IDENTIFIER "z" is of TYPE int
// ASSIGNS the ZERO VALUE of TYPE int to "z"
var z int

func main() {
	// short declaration operator
	//DECLARE a varibale and ASSIGN a VALUE (of a certain TYPE)
	x := 42
	fmt.Println(x)

	foo()

	fmt.Println(z)
}

func foo() {
	fmt.Println("again: ", y)
}

// x := [] 숏컷은 함수 안에서만 가능하다.
// var y = [] 는 함수 밖에서도 가능하다.
