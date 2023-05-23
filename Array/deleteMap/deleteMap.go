package main

import (
	"fmt"
)


func main() {
	m := map[string]int{
		"James": 32,
		"Miss MoneyPenny": 27,
	}

	m["todd"] = 33

	for k, v := range m {
		fmt.Println(k, v)
	}

	xi := []int{1,2,3,4,5,6,7,8,9}
	for i, v := range xi {
		fmt.Println(i, v)
	}

	delete(m, "James")
	fmt.Println(m)
	fmt.Println(m["Ian Fleming"])

	if v, ok := m["Miss MoneyPenny"]; ok {
		fmt.Println("value:", v)
		delete(m, "Miss MoneyPenny")
	}
}