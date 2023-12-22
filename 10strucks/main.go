package main

import "fmt"

func main() {
	fmt.Println("Example of Struck")

	saty := User{"Satyendra", "satyendradhamgaye2@gmail.com", true, 20}
	//this unstructured way of printing a struck
	fmt.Println(saty)
	//this is a structured wy of printing a struck
	fmt.Printf("Details of Satyendra are:\n%+v\n", saty)
	//for printing perticular values
	fmt.Printf("Name is: %v, Email is: %v", saty.Name, saty.Email)
}

type User struct {
	Name  string
	Email string
	Auth  bool
	Age   int
}
