package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("we got an error")
	}
	return a / b, nil
}
func main() {
	res, err := divide(10, 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Result", res)
}
