package main

import "fmt"

func describe(i interface{}) {

	fmt.Printf("type %T and value %v\n", i, i)
}

func main() {
	describe(10)
	describe("Amaresha")
	describe(12.8)
}
