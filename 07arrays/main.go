package main

import "fmt"

func main() {
	fmt.Println("Arrays")

	var carList [3]string

	carList[0] = "BMW"
	carList[1] = "AUDI"
	carList[2] = "TOYOTA"

	fmt.Println("Values insisde carlist are: ", carList)
	fmt.Println("Lenght of the Array carlist is: ", len(carList))
	fmt.Printf("The type of CarList is %T\n", carList)

	var number = [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println("Values insisde carlist are: ", number)
	fmt.Println("Lenght of the Array carlist is: ", len(number))
}
