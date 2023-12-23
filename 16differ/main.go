package main

import "fmt"

func main() {
	//defer is keyword that we user in front of a statement
	//when used, this make the statement to execute at the very bottom of the function

	defer fmt.Println("Statement Two")
	fmt.Println("Statement One")

	//now if there are multiple difer keyword one after another
	//the flow will of LIFO(Last In First Out)

	defer fmt.Println("Statement Three")
	defer fmt.Println("Statement Four")
	defer fmt.Println("\nStatement Five")

	//now if there is a function which contains a defer value
	//then it'll execute first then the remaining defers lets see an example
	myDefer()
}

func myDefer() {
	for i := 0; i < 4; i++ {
		fmt.Print(i)
	}
}
