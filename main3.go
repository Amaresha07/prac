package main

import "fmt"

type user struct {
	name string
	age  int
}

func (p user) greet(village string, dist string) {
	fmt.Printf("My name is %v and age is %v and my village is %v and my dist is %v\n", p.name, p.age, village, dist)
}

func main() {
	p := user{"Amaresha", 21}
	p.greet("uppala", "raichur")

}
