package main

import (
	"fmt"
	"sort"
)

func main() {
	var carList = []string{"AUDI", "Tesla"}

	//using append in slice to add more values in it
	carList = append(carList, "BMW", "Toyota")
	fmt.Println(carList)

	//here can use the below syntax to start a slice and end this where ever
	//it is required
	carList = append(carList[1:3])
	fmt.Println(carList)

	numbers := make([]int, 3)

	numbers[0] = 103413
	numbers[1] = 205436
	numbers[2] = 304524
	fmt.Println("Length before append is: ", len(numbers))
	fmt.Println("Values inside the slice is: ", numbers)
	numbers = append(numbers, 405432, 54534, 65450)
	fmt.Println("Length after append is: ", len(numbers))
	fmt.Println("Values inside the slice is: ", numbers)

	fmt.Println("If the list inside the slice numbers are sorted or not: ", sort.IntsAreSorted(numbers))

	sort.Ints(numbers)
	fmt.Println("Latest sorted number list: ", numbers)
	fmt.Println("If the list inside the slice numbers are sorted or not: ", sort.IntsAreSorted(numbers))

	//lets remove a value from a list via append function

	var tech = []string{"LG", "Lenove", "DELL", "Apple", "Google"}
	fmt.Println("Current list of the tech slice is: ", tech)
	index := 1
	tech = append(tech[:index], tech[index+1:]...)
	fmt.Println("tech slice list after removing Lenove form the list: ", tech)

}
