package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	//Positive or negetive Numebr
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter a value to determined wather it is positvie or negetive: ")
	value, _ := reader.ReadString('\n')
	number, _ := strconv.ParseInt(strings.TrimSpace(value), 10, 64)

	if number > 0 {
		fmt.Println("Value entered is positive")
	} else if number < 0 {
		fmt.Println("Value entered is negetive")
	} else {
		fmt.Println("Value entered is Zero")
	}

	//Even or odd number
	fmt.Println("\nIf the value entered is Even or Odd: ")
	if number%2 == 0 {
		fmt.Println("Value entered is even")
	} else {
		fmt.Println("Values entered is odd")
	}

	//assigning a varibale in the if conditon

	fmt.Println("\nIf the variable 6 is greater or smaller then 10: ")
	if num := 6; num > 10 {
		fmt.Println("Greater then 10")
	} else {
		fmt.Println("Smaller then 10")
	}
}
