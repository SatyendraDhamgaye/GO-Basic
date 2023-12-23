package main

import "fmt"

//createing calulator using function only
func main() {
	a := 10.156
	b := 2.567

	fmt.Println("Given number are 10.156 and 2.567")
	fmt.Printf("Additon of given number is: %.2f\n", add(a, b))
	fmt.Printf("Substraction of given number is: %.2f\n", sub(a, b))
	fmt.Printf("Multiplication of given number is: %.2f\n", mul(a, b))
	fmt.Printf("Division of given number is: %.2f\n", div(a, b))

	//addition of a slice
	fmt.Println("Total of n numbers are: ", totalOfSlice(1, 2, 3, 4, 5, 6, 7))
}

func add(val1 float64, val2 float64) float64 {
	return val1 + val2
}

func sub(val1 float64, val2 float64) float64 {
	return val1 - val2
}

func mul(val1 float64, val2 float64) float64 {
	return val1 * val2
}

func div(val1 float64, val2 float64) float64 {
	return val1 / val2
}

func totalOfSlice(total ...int) int {
	value := 0
	for _, val := range total {
		value += val
	}
	return value
}
