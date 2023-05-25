package main

import "fmt"

func main() {
	xi := []int{2, 3, 4, 5, 6, 7 ,8 ,9}
	x := sum("James Bond", xi...)
	fmt.Println(x)

}


func sum(s string, x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0

	for _, v := range x {
		sum += v
	}

	fmt.Println("Total is:", sum)
	return sum
}