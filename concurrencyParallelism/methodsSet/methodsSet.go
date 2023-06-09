package main

import "fmt"

type circle struct   {
	radius float64
}

type shape interface {
	area() float64
}

func (c *circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area: ", s.area())
}

func main() {
	c := circle{5}
	info(&c)
}