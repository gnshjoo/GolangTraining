package main

import "fmt"

func main() {
	x := []int{4, 2, 1, 34, 23}
	fmt.Println(x)
	fmt.Println(x[1:])
	fmt.Println(x[1:3])

	for i, v := range x {
		fmt.Println(i, v)
	}

	for i := 0; i<=5; i++ {
		fmt.Println(i, x[i])
	}
}

