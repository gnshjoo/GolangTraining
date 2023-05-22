package main

import "fmt"

// Using iota, create 4 constants for the NEXT 4 years. Print the constant values.

const (
	a1 = 2021 + iota
	b1
	c
	d
	e
)

func main() {
	fmt.Println(a1, b1, c, d, e)
}
