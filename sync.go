package main

import (
	"fmt"
	"sync"
	"time"
)

func processorders(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Processing order %v\n", i)
	time.Sleep(time.Second)
	fmt.Printf("order  %v completed\n", i)

}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go processorders(i, &wg)
	}
	wg.Wait()
	fmt.Println("All oreders are completed")
}
