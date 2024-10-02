package main

import "fmt"

func main() {

	var ptr *int

	fmt.Println("Value of pointer is", ptr) // nil

	var num int = 10

	ptr = &num // address of num
	fmt.Println("Value of pointer is", ptr) // address of num
	fmt.Println("Value of num is", *ptr) // actual value of num (10 here)

	*ptr *= 3 // manipulating the actual value of num using the pointer 

	fmt.Println("Value of num is", num) // 30
}
