package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a) // 주소값이 출력된다.

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Println(b)
	fmt.Println(*b) // *를 붙이면 주소값이 아닌 값이 출력된다.

	fmt.Println(*&a)

	*b = 43
	fmt.Println(a)

}

// go의 값은 언제나 값으로 전달된다.