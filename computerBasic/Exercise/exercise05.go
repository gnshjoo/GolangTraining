package main

import "fmt"

// Create a variable of type string using a raw string literal. Print it
func main() {
	a := `This is a raw string literal
	"you see" 
	another line`
	fmt.Println(a)
}
