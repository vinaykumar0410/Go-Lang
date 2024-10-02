package main

import (
	"fmt"
	"math/rand"
)

func main() {

	diceNumber := rand.Intn(6) + 1

	fmt.Println("Dice number is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice number is One")
	case 2:
		fmt.Println("Dice number is Two")
		// fallthrough // used to print next case along with current case 
	case 3:
		fmt.Println("Dice number is Three")
	case 4:
		fmt.Println("Dice number is Four")
	case 5:
		fmt.Println("Dice number is Five")
	case 6:
		fmt.Println("Dice number is Six")
	}


}
