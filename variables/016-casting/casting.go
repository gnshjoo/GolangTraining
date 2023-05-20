package main

import "fmt"

var a int

type hotdog int

var b hotdog

func main() {
	a = 42
	fmt.Println(a)
	fmt.Println("%T\n", a)

	fmt.Println(b)
	fmt.Println("%T\n", b)
	a = int(b)
	fmt.Println(a)
	fmt.Println("%T\n", a)
}
