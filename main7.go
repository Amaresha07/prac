package main

import "fmt"

type Address struct {
	city  string
	state string
}

type employee struct {
	name    string
	age     int
	address Address
}

func main() {
	p := employee{name: "Amaresha", age: 21, address: Address{city: "Raichur", state: "karnataka"}}
	fmt.Printf("%v and is %v and is from %v and is state is %v\n", p.name, p.age, p.address.city, p.address.state)
}
