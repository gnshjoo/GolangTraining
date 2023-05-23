package main

import "fmt"

func main() {
	x := make([]int, 10, 12)
	fmt.Println(
		x,
		len(x),
		cap(x),
	)

	x[0] = 42
	x[9] = 999

	x = append(x, 3234)

	fmt.Println(x, len(x), cap(x))

	x = append(x, 3234)
	x = append(x, 3234)

	fmt.Println(x, len(x), cap(x))
}