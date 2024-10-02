package main

import (
	"fmt"
	"sync"
)

func main() {

	// init waitgroup and channel 
	wg := &sync.WaitGroup{}

	ch := make(chan int)

	// add 2 go routines to waitgroup 
	wg.Add(3)
	// Send data to channel 
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// send data to channel
		ch <- 3
		ch <- 5
	}(wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// read data from channel (no problem even if channel is closed)
		// sending data should match with receiving data
		fmt.Println(<-ch) 
		fmt.Println(<-ch)
	}(wg)


	// send data to channel
	go sendData(ch,wg)

	// read data from channel
	for val := range ch {
		fmt.Println("Received:", val)
	}

	wg.Wait()
}


func sendData(ch chan int, wg *sync.WaitGroup) {

	defer wg.Done()

	// send data to channel
	for i := 1; i < 6; i++ {
		fmt.Println("Sending", i)
		ch <- i 
	}

	// close channel to indicate that no more data will be sent
	close(ch) 
}
