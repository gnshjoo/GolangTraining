package main

import "fmt"

// Create your own type. Have the underlying type be an int.
// create a VARIABLE of your new TYPE with the IDENTIFIER "x" using the "VAR" keyword
// in fuc main
//
//	a. print out the value of the variable "x"
//	b. print out the type of the variable "x"
//	c. assign 42 to the VARIABLE "x" using the "=" OPERATOR
//	d. print out the value of the variable "x"
type customer int

var x2 customer

func main() {
	fmt.Println(x2)
	fmt.Printf("%T\n", x2)
	x2 = 42
	fmt.Println(x2)
}
