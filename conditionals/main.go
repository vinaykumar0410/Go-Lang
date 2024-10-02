package main

import "fmt"

func main() {

	var age int = 23

	if age >= 18 {
		fmt.Println("You are eligible to vote")
	} else {
		fmt.Println("You are not eligible to vote")
	}

	// You are eligible to vote

	var a int = 20
	var b int = 30
	var c int = 10

	if a > b && a > c {
		fmt.Println("a is greater than b & c")
	} else if b > c {
		fmt.Println("b is greater than a & c")
	} else {
		fmt.Println("c is greater than a & b")
	}

	// b is greater than a & c

	if number := 40; number == 40 {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not Equal")
	}
	
	// Equal

}