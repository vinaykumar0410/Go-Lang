package main

import "fmt"

func main() {

	// defer will be executed in LIFO order

	defer fmt.Println("Happy New Year!!")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")

	fmt.Println("Count Down Started!!")

	reverseFromNtoOne(5)
	
	fmt.Println("\nDone With Normal Execution!! Now Executing Defer")
	/*
		Count Down Started!!
		5 4 3 2 1 
		Done With Normal Execution!! Now Executing Defer
		Three
		Two
		One
		Happy New Year!!
	*/
}


func reverseFromNtoOne(num int){
	for i := 1; i <= num; i++ {
		defer fmt.Printf("%d ", i)
	}
}