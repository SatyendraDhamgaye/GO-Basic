package main

import "fmt"

func main() {
	saty := User{"Satyenrda", "satyendra@gmail.com", true, "Online", 21, 82.9}
	saty.GetAuth()
	saty.NewMail()
	fmt.Println(saty.Email)
}

type User struct {
	Name       string
	Email      string
	Auth       bool
	Status     string
	Age        int
	Percentage float64
}

func (u User) GetAuth() {
	fmt.Println("Is the user Authorized: ", u.Auth)
}

func (u User) NewMail() {
	u.Email = "new@mail.com"
	fmt.Println(u.Email)
}
