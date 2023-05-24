package main

import "fmt"

// we create VALUES of a certain TYPE that are stored in VARIABLES
// and those VARIABLES have IDENTIFIERS

var x int

type person struct {
	first string
	last string
}

type foo int
var y foo


func main() { 
	p1 := person{
		first: "James",
		last: "Bond",
	}
	fmt.Println("Hello, playground", p1)
}