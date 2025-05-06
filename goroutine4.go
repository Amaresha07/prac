package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	//fmt.Printf("worker %d completed\n", id)
	time.Sleep(1 * time.Millisecond)
	ch <- fmt.Sprintf("worker %d completed\n", id)
}

func main() {
	var wg sync.WaitGroup

	ch := make(chan string, 2)

	wg.Add(2)

	go worker(1, &wg, ch)
	go worker(2, &wg, ch)

	wg.Wait()
	close(ch)
	for mesg := range ch {
		fmt.Println(mesg)
	}
	fmt.Println("All goroutne finished")
}
