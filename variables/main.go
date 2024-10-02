package main

import "fmt"

const PublicVaiable string = "Why Public Bcoz 1st Letter is Capital"

func main() {

	var username string = "Vinay"
	fmt.Printf("%s is of type : %T \n", username, username)

	var age int = 23
	fmt.Printf("%d is of type : %T \n", age, age)

	var isStudent bool = false
	fmt.Printf("%v is of type : %T \n", isStudent, isStudent)

	const pi float64 = 3.14159
	fmt.Printf("%f is of type : %T \n", pi, pi)

	var complexNum complex128 = 2 + 3i
	fmt.Printf("%v is of type : %T \n", complexNum, complexNum)

	var runeValue rune = 'A'
	fmt.Printf("%c is of type : %T \n", runeValue, runeValue)

	var fruit = "apple"
	// fruit = 20 // error 
	fmt.Println(fruit)
	fruit = "mongo" 
	fmt.Println(fruit)

	// allow only inside method
	noOfFriends := 10
	fmt.Println(noOfFriends)

	fmt.Println(PublicVaiable)
}