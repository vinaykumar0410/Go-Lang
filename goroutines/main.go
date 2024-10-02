package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

var mutux sync.Mutex

var workingEndpoints []string

func main() {

	// go greet("Vinay")
	// greet("Sai")
	_ = greet

	var endpoints = []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.stackoverflow.com",
		"https://vinay-nextjs.vercel.app",
		"https://www.twiter.com", // invalid end point
	}

	for _, endpoint := range endpoints {
		go getStatusCode(endpoint)
		wg.Add(1)
	}
		
	wg.Wait()
	
	// valid end points
	fmt.Println("Valid End Points are", workingEndpoints)
}

func greet(name string) {

	for i := 1; i < 5; i++ {
		
		fmt.Println("Hello", name)
		time.Sleep(2 * time.Second)

	}

}


func getStatusCode(endpoint string) {
	
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Printf("Issue with %s endpoint\n", endpoint)
	} else {
		
		mutux.Lock()
		workingEndpoints = append(workingEndpoints, endpoint)
		mutux.Unlock()

		fmt.Printf("%d StatusCode for %s endpoint\n", res.StatusCode, endpoint)
	}

}
