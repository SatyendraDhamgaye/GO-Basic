package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a value")

	// comma ok // comma err
	// Syntax: variable, error := reader.Read('\n') where variable is the value and error if someting goes wrong

	value, _ := reader.ReadString('\n')
	fmt.Print("Entered value is ", value)
	fmt.Printf("The type of the enterd value is %T", value)

}
