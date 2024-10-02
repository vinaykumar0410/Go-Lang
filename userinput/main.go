package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("Enter your age : ")

	// read input from user
	reader := bufio.NewReader(os.Stdin)
	// comma ok || error ok syntax
	input, _ := reader.ReadString('\n')

	fmt.Print("Your age is : ",input)

	// wish to add 1 more to age as today is my birthday
	convertedInput, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)
	fmt.Println("Your age is now :",convertedInput + 1)
}