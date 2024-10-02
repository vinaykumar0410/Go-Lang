package main

import (
	"fmt"
	"sort"
)

func main() {

	var fruits = []string{"apple", "banana", "mango"}

	fmt.Println("Fruits are : ", fruits) // Fruits are :  [apple banana mango]

	fruits = append(fruits, "pineapple", "grapes") 

	fmt.Println("Modified Fruits are : ", fruits) // Modified Fruits are :  [apple banana mango pineapple grapes]

	fmt.Println("Length of fruits is : ", len(fruits)) // Length of fruits is :  5
	
	fruits[len(fruits)-1] = "orange"
	fmt.Println("Modified fruits are :", fruits) // Modified fruits are :  [apple banana mango pineapple orange]

	// i dont want mango
	var index int = 2 
	fruits = append(fruits[:index], fruits[index+1:]...)

	for fruit := range fruits {
		fmt.Println("Fruit is :", fruits[fruit]) // prints all fruits expects mango
	}

	numbers := make([]int, 3)
	fmt.Println("Numbers are : ", numbers) // Numbers are :  [0 0 0]

	numbers[0] = 8;
	numbers[1] = 4;
	numbers[2] = 5;

	fmt.Println("Numbers are : ", numbers) // Numbers are :  [8 4 5]
	
	// numbers[4] = 2; // panic: runtime error: index out of range [4] with length 3

	numbers = append(numbers, 2, 7, 1)

	fmt.Println("Numbers are : ", numbers) // Numbers are :  [8 4 5 2 7 1]
	
	fmt.Println("Is sorted :",sort.IntsAreSorted(numbers)) // false
	
	sort.Ints(numbers)
	
	fmt.Println("Numbers are : ", numbers) // Numbers are :  [1 2 4 5 7 8]
	fmt.Println("Is sorted :",sort.IntsAreSorted(numbers)) // true
	

}
