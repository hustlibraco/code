package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func printer(ch chan int) {
	for i := range ch {
		fmt.Printf("Received %d ", i)
	}
	time.Sleep(3 * time.Second)
	fmt.Println("All Done")
	wg.Done()
}

// main is the entry point for the program.
func main() {
	c := make(chan int)
	go printer(c)
	wg.Add(1)

	// Send 10 integers on the channel.
	for i := 1; i <= 10; i++ {
		c <- i
	}

	close(c)
	wg.Wait()
}
