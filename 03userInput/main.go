package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a value")
	value, _ := reader.ReadString('\n')
	fmt.Print("Entered value is ", value)
	fmt.Printf("The type of the enterd value is %T", value)

}
