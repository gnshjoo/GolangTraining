package main

import "fmt"

var y = 42

// DECLARE the VARIABLE with the IDENTIFIER "z"
// is of TYPE string
// and I ASSIGN the VALUE "Shaken, not stirred"
var z = "Shaken, not stirred"

// Thjs is a STATIC programming language
// a VARIABLE is DECLARED to hold a VALUE of a certain TYPE
// no a DYNAMIC programming language

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
}
