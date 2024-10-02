package main

import (
	"fmt"

)

func main() {

	greet()

	var num1 int = 20
	var num2 int = 10

	maxNum, minNum := maxmin(num1,num2)

	fmt.Println("Max number is",maxNum)
	fmt.Println("Min number is",minNum)

	var elements = []int {3,4,1,5,6,2}

	result := sum(elements...)

	fmt.Println("Sum of Elements is :",result)

	func (){
		fmt.Println("Immediate Calling Function You Can Call It As Anonymous Function")
	}()
}

func greet() {
	fmt.Println("Hey How Are You")

}

func maxmin(num1 int,num2 int) (int, int){
	if num1 > num2 {
		return num1, num2
	}
	return num2, num1
}

func sum(numbers ...int) int {

	var total int = 0

	for _, num := range numbers {
		total += num
	}

	return total
}