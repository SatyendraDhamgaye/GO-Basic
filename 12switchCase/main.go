package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("D6 dice game: ")

	rand.Seed(time.Now().UnixNano())
	diceNum := rand.Intn(6) + 1

	switch diceNum {
	case 1:
		fmt.Println("Dice rolled a 1, can open a peace")
	case 2:
		fmt.Println("Dice rolled a 2")
	case 3:
		fmt.Println("Dice rolled a 3")
	case 4:
		fmt.Println("Dice rolled a 4")
	case 5:
		fmt.Println("Dice rolled a 5")
	case 6:
		fmt.Println("Dice rolled a 6, can roll again and open a peace")
	default:
		fmt.Println("YOU CHEATER!!!")
	}
}
