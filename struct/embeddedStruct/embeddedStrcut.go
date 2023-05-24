package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

type seceretAgent struct {
	person
	ltk bool
}

func main() {
	p1 := seceretAgent{
		person: person{
			first: "James",
			last: "Bond",
			age: 32,
		},
		ltk: true,


	}
	fmt.Println(p1.first, p1.last, p1.age, p1.ltk)
}

// inner type의 outter type 에 같은 이름이 없을 경우 생략가능하다.