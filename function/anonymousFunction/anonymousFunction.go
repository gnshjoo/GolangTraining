package main

import "fmt"

func main() {
	func() {
		fmt.Println("I am an anonymous function")
	}()

	func (x int) {
		fmt.Println("The meaning of life:", x)
	}(42)

	f := func() {
		fmt.Println("My first func expression")
	}

	f()
}