package main

import "fmt"

func main() {
	// for init; condition; post {}
	for i := 0; i <= 100; i++ {
		for j := 0; j <= 100; j++ {
			fmt.Println("i:", i, "j:", j)
			
		}
	}
}