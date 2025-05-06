package main

import "fmt"

type rectangle struct {
	length, breadth float64
}

func (r rectangle) area() float64 {
	return r.length * r.breadth
}

func main() {
	r := rectangle{2.7, 4.8}
	fmt.Println("the area of rectangle", r.area())
}
