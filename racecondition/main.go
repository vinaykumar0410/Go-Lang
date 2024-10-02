package main

import (
	"fmt"
	"sync"
)

func main() {

	// init numbers array with default value "Nil"
	var numbers = []string{"Nil"}

	// init waitgroup
	wg := &sync.WaitGroup{}

	// init mutex
	mutux := &sync.Mutex{}

	// init read-write mutex
	mutuxRW := &sync.RWMutex{}

	// can run 4 go routines in one go
	wg.Add(4)
	// wg.Add(1) // run 1 go routine
	go func(mut *sync.Mutex) {
		fmt.Println("Zero")
		// lock the mutex to prevent race condition
		mutux.Lock()
		numbers = append(numbers, "Zero")
		// unlock the mutex
		mutux.Unlock()
		// release the waitgroup
		wg.Done()
	}(mutux)

	// wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("One")
		mut.Lock()
		numbers = append(numbers, "One")
		mut.Unlock()
		wg.Done()
	}(wg, mutuxRW)

	// wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Two")
		mut.Lock()
		numbers = append(numbers, "Two")
		mut.Unlock()
		wg.Done()
	}(wg, mutuxRW)

	// wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Three")
		mut.Lock()
		numbers = append(numbers, "Three")
		mut.Unlock()
		wg.Done()
	}(wg, mutuxRW)

	// wait for all go routines to finish
	wg.Wait()

	// lock the read-write mutex to prevent race condition
	mutuxRW.RLock()
	fmt.Println(numbers)
	mutuxRW.RUnlock()

}
