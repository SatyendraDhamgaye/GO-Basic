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

	fmt.Println("Enter a vaule for a pointer example")
	numberinstr, _ := reader.ReadString('\n')
	number, _ := strconv.ParseInt(strings.TrimSpace(numberinstr), 10, 64)
	fmt.Println("Entered number is :", number)

	//here & means the address of the variables that stored inside
	ptr := &number
	fmt.Println("pointer address for the values you enterd is: ", ptr)

	//here * means the what data is stored inside that vointer not the actual address like &
	fmt.Println("values that is inside of the pointer is: ", *ptr)

}
