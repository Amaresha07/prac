package main

import "fmt"

func main() {

	var i interface{} = 13.8

	str, ok := i.(float64)

	if ok {
		fmt.Println("vakue", str)
	} else {
		fmt.Println("type asssertion failed")
	}
}
