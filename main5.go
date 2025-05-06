package main

import "fmt"

type counter struct {
	count int
}

func (c *counter) increment() {
	c.count--
}

func main() {
	c := counter{5}
	c.increment()
	fmt.Println(c.count)
}
