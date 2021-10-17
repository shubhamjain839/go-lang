package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to switches")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	fmt.Println("Number is: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is one")

	case 2:
		fmt.Println("Dice value is two")
	case 3:
		fmt.Println("Dice value is three")
	case 4:
		fmt.Println("Dice value is four")
	case 5:
		fmt.Println("Dice value is five")
	case 6:
		fmt.Println("Dice value is six")
	default:
		fmt.Println("invalid number-+")
	}
}
