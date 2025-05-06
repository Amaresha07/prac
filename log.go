// package main

// import "log"

// func main() {
// 	log.Println("server is started")
// 	log.Println("server is giving warning")
// 	log.Println("server is stoped")

// }

package main

import "fmt"

func main() {
	ch := make(chan int, 3) // Channel buffer size is 2
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println(<-ch) // Prints 1
	fmt.Println(<-ch) // Prints 2
	fmt.Println(<-ch)
}
