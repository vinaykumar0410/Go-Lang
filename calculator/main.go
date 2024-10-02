package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("CALCULATOR USING GO LANG")

	reader := bufio.NewReader(os.Stdin)

	for {

		// input-1
		fmt.Printf("Enter a number: ")
		number1, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Invalid Input Try Again!!")
			continue
		}
		number1 = strings.TrimSpace(number1)
		if number1 == "exit" {
			break
		}

		num1, err := strconv.Atoi(number1)
		if err != nil {
			fmt.Println("Invalid Input Try Again!!")
			continue
		}

		// input-2
		fmt.Printf("Enter another number: ")
		number2, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Invalid Input Try Again!!")
			continue
		}
		number2 = strings.TrimSpace(number2)
		if number2 == "exit" {
			break
		}

		num2, err := strconv.Atoi(number2)
		if err != nil {
			fmt.Println("Invalid Input Try Again!!")
			continue
		}

		// operator
		fmt.Println("Enter operator from these list [e + - * / % ] e for exit : ")
		operator, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Invalid Input Try Again!!")
			continue
		}
		operator = strings.TrimSpace(operator)
		if operator == "e" {
			break
		}

		switch operator {
			case "+":
				fmt.Println("Sum of two numbers:", num1+num2)
			case "-":
				fmt.Println("Difference of two numbers:", num1-num2)
			case "*":
				fmt.Println("Product of two numbers:", num1*num2)
			case "/":
				if num2 == 0 {
					fmt.Println("Zero Division Error")
					continue
				}
				fmt.Println("Quotient of two numbers:", num1/num2)
			case "%":
				fmt.Println("Remainder of two numbers:", num1%num2)
			default:
				fmt.Println("Invalid Operator Try Again!!")
		}
	}

	fmt.Println("Exiting the calculator. Goodbye!")
}
