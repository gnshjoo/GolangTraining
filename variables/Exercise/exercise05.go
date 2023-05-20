package main

import "fmt"

// Building on the code from the previous example
// 1. at the package level scope, using the “var” keyword, create a VARIABLE with the
// IDENTIFIER “y”. The variable should be of the UNDERLYING TYPE of your custom TYPE “x”
// 2. in func main
//	a. this should already be done
//		i. print out the value of the variable “x”
//		ii. print out the type of the variable “x”
//		iii. assign your own VALUE to the VARIABLE “x” using the “=” OPERATOR
//		iv. print out the value of the variable “x”
//	b. now do this
//		i. now use CONVERSION to convert the TYPE of the VALUE stored in “x” to the
//		UNDERLYING TYPE
//			1. then use the “=” operator to ASSIGN that value to “y”
//			2. print out the value stored in “y”

type hotdog int

var x4 hotdog
var y4 int

func main() {
	fmt.Println(x)
	fmt.Println("%T\n", x4)
	x4 = 42
	fmt.Println(x4)
	y4 = int(x4)
	fmt.Println(y4)
	fmt.Println("%T\n", y4)

}
