package main

import "fmt"

//global variables
var globalvar uint = 10

//constant variavle
const constantvar float64 = 538.452

func main() {

	//default variables

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

	//default variables with no values

	var nothing string
	fmt.Println(nothing)
	fmt.Printf("The above value is of %T type \n", nothing)

	//with no datatype

	var notype = "abc"
	fmt.Println(notype)
	fmt.Printf("The above value is of %T type \n", notype)

	//no var type direct approch

	direct := 10.5
	fmt.Println(direct)
	fmt.Printf("The above value is of %T type \n", direct)

	fmt.Println(globalvar)
	fmt.Printf("The above value is of %T type \n", globalvar)

	fmt.Println(constantvar)
	fmt.Printf("The above value is of %T type \n", constantvar)
}
