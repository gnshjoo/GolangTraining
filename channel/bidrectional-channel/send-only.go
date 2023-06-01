package main

import "fmt"

func main() {
	c := make(chan<- int, 2)

	c <- 42
	c <- 53 // send only

	fmt.Println("-----")
	fmt.Printf("%T\n", c)
}
