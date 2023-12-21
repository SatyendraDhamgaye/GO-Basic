package main

import "fmt"

func main() {
	var name string = "David"
	fmt.Println(name)
	fmt.Printf("The above value is of %T type \n", name)

	var pvalue uint = 10
	fmt.Println(pvalue)
	fmt.Printf("The above value is of %T type \n", pvalue)

	var nvalue int = -20
	fmt.Println(nvalue)
	fmt.Printf("The above value is of %T type \n", nvalue)

	var fvalue float32 = 10.24
	fmt.Println(fvalue)
	fmt.Printf("The above value is of %T type \n", fvalue)

	var validation bool = false
	fmt.Println(validation)
	fmt.Printf("The above value is of %T type \n", validation)
}
