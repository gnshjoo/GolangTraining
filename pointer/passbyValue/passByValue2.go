package main

import "fmt"

func main2() {
	x := 0
	fmt.Println("x before", x)
	fmt.Println("x before", &x)
	foo2(&x)
}

func foo2(y *int) {
	fmt.Println("y before", y)
	fmt.Println("y before", *y)
	*y = 43
	fmt.Println("y after", y)
	fmt.Println("y after", *y)
}