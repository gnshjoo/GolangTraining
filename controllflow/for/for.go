package main

import "fmt"

func main() {
	// for statement single codition
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("Single condititon Done.")

	// for statement with for clause
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}

	fmt.Println("for clause Done.")

	// for statements with range clause

	slice := []string{"a", "b", "c"}
	for i, v := range slice { 
		fmt.Println(i, v)
	}
	fmt.Println("range clause Done.")

}