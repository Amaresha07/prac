package main

import (
	"fmt"
)

// Define an interface
type Animal interface {
	Speak()
}

// Dog struct
type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("Woof! Woof!")
}

// Cat struct
type Cat struct{}

func (c Cat) Speak() {
	fmt.Println("Meow!")
}

type tiger struct{}

func (t tiger) Speak() {
	fmt.Println("Goorrrr")
}

func makeSound(a Animal) {
	a.Speak() // Works for any type that implements Speak()
}

func main() {
	d := Dog{}
	c := Cat{}
	t := tiger{}

	makeSound(d) // Woof! Woof!
	makeSound(c) // Meow!
	makeSound(t)
}
