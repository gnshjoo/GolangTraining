package main

import (
	"fmt"
)

func main() {
	switch  {
	case false:
		println("this should not print")
	case (2==4):
		fmt.Println("this should not print2")
	case 3 != 3:
		fmt.Println("prints")
		//fallthrough
	case 4 != 4:
		fmt.Println("also true, does it print?")
		//fallthrough
	case 7 == 4:
		fmt.Println("also true, does it print?")
		//fallthrough
	case 11 == 4:
		fmt.Println("also true, does it print?")
	default:
		fmt.Println("this is default")
	}
}