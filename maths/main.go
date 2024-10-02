package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	// "math/rand"
)

func main() {

	var bagOneCoins int = 1350
	var bagTwoCoins float64 = 1600.50

	fmt.Println("Total Coins from both bags :", bagOneCoins + int(bagTwoCoins)) // precision loss

	// generating random number using math module

	// var randomNumber int = rand.Intn(10) + 1

	// fmt.Println("Random number is", randomNumber)

	luckyNumber, err := rand.Int(rand.Reader, big.NewInt(6))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Lucky Number is", luckyNumber)
}
