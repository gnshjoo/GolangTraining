package main

import "fmt"

func main() {
	fmt.Println(mySum(2, 3))
	fmt.Println(mySum(3, 55))
	fmt.Println(mySum(12, 2))
}

func mySum(xi ...int) int {
	var sum int
	for _, v := range xi {
		sum += v
	}
	return sum
}
