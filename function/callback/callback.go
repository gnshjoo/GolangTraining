// package main

// import "fmt"

// func main() {
// 	x := 0
// 	fmt.Println(x)
// 	x++
// 	fmt.Println(x)

// 	x += 42
// 	fmt.Println(x)

// 	x-=3
// 	fmt.Println(x)

// 	x *= 2
// 	fmt.Println(x)

// 	x--
// 	fmt.Println(x)
// }

package main

import "fmt"

func main() {
	ii := []int{1,2,3,4,5,6,7,8,9,10}
	s := sum(ii...)
	fmt.Println(s)

	e := odd(sum,ii...)
	fmt.Println(e)
}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}

	return total
}

func odd(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v % 2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}