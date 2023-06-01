package main

import "fmt"

func main() {
	c := make(<-chan int, 2)

	fmt.Println(<-c)
	fmt.Println(<-c) // receive only

	fmt.Println("-----")
	fmt.Printf("%T\n", c)
}
