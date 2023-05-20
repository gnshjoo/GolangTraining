package main

import "fmt"

// Use var to DECLARE three variables. The variables should have package level scope.
// Do not assign VALUES to the variables. Use the following IDENTIFIERS for the variables
// and make sure the variables are of the following TYPE.
// a. identifier “x” type int
// b. identifier “y” type string
// c. identifier “z” type bool
// in func main
// a. print out the values for each identifier
// b. The compiler assigned values to the variables. What are these values called? (Zero values)

var x int
var y string
var z bool

func main() {
	fmt.Println(x, y, z)
}
