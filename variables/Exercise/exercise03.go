package main

import "fmt"

// Using the code form the previous exercise,
// 1. At the package level scope, assign the following values to the three variables
//  a. for x assign 42
//	b. for y assign “James Bond”
//	c. for z assign true
// 2. in func main
//  a. use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the
//  returned value of TYPE string using the short declaration operator to a VARIABLE
// 	with the IDENTIFIER “s”
//  b. print out the value stored by variable “s”

var x1 = 42
var y1 = "James Bond"
var z1 = true

func main() {
	s := fmt.Sprintf("%v\t%v\t%v", x1, y1, z1)
	fmt.Println(s)
}
