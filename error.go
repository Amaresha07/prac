package main

import "fmt"

func main() {
	var num int

	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("we got an error")
		fmt.Println(err)
	} else {
		fmt.Println(num)
	}
}
