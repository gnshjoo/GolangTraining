package main

import "fmt"

func main() {
	foo()
	bar("James")
	s1 :=woo("Moneypenny")
	fmt.Println(s1)
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x, y)
}

// func (r reciver) identifier(parameters) (return(s)) { ... }
// know the diffrence between parameter and argument
//  - we defined our func with parameters (if any)
//  - we call our func and pass in arguments (if any)

func foo() {
	fmt.Println("Hello from foo")
}

// everthing in Go is PASS BY VALUE
func bar(s string ){
	fmt.Println("Hello,", s)
}

func woo(st string) string{
	return fmt.Sprint("Hello from woo, ", st)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, `, says "Hello"`)
	b := false
	return a, b
}