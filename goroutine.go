package main

import (
	"fmt"
	"time"
)

func printmesg() {
	fmt.Println("hello fron gorotunie")
}

func main() {
	go printmesg()

	fmt.Println("hello from main")
	time.Sleep(time.Second)
}
