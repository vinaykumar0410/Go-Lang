package main

import "fmt"

func main() {

	var technologies = []string{"Java", "Go", "Python", "C++"}

	for i := 0; i < len(technologies); i++ {
		fmt.Printf(technologies[i] + " ") // Java Go Python C++
	}

	fmt.Println("")
	
	for i := range technologies {
		fmt.Printf(technologies[i] + " ") // Java Go Python C++
	}
	
	fmt.Println("")

	for index, technology := range technologies { // can use _, technology instead
		fmt.Println(index, "->", technology) // 0 -> Java 1 -> Go 2 -> Python 3 -> C++
	}

	breakPoint := 6

	num := 1

	for num < 10 {
		if num == breakPoint {
			break // breaks out of the loop
		}
		fmt.Printf("%d ", num) // 1 2 3 4 5
		num++
	}

	fmt.Println("")

	fmt.Println(num) // 6

	blackDay := 8

	birthDay := 9

	for num < 10 {
		if num == birthDay {
			goto greetings // jumps to the label
		}
		if num == blackDay {
			num++
			continue // continues the loop
		}
		fmt.Printf("%d ", num) 
		num++
	}

	greetings:
		fmt.Printf("Happy Birthday! %d is my favourite number too", num)

	// 6 7 Happy Birthday! 9 is my favourite number too

}