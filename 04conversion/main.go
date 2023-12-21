package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Addition of numbers")
	fmt.Println("Enter number 1: ")
	input1, _ := reader.ReadString('\n')
	fmt.Println("Enter number 1: ")
	input2, _ := reader.ReadString('\n')

	number1, err1 := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	number2, err2 := strconv.ParseFloat(strings.TrimSpace(input2), 64)

	//	number1, err1 := strconv.ParseInt(strings.TrimSpace(input1), 10, 64)
	//  number2, err2 := strconv.ParseInt(strings.TrimSpace(input2), 10, 64)

	if err1 != nil && err2 != nil {
		fmt.Println("error is :", err1, err2)
	} else {
		sum := number1 + number2
		fmt.Println("addition of two numbers is", sum)
	}

}
