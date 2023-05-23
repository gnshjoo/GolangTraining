package main

import "fmt"

func main() {
	m := map[string]int{
		"James": 32,
		"Miss MoneyPenny": 27,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Barnabas"]) //없는 key 값을 입력하면 0을 반환한다.

	v, ok := m["Barnabas"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Barnabas"]; ok {
		fmt.Println("THIS IS THE IF PRINT", v)
	}

	
}