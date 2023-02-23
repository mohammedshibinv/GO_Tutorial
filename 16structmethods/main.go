package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
}

func main() {
	var newUser = User{"Shibin", "Shibin@gamil.com", true}
	newUser.getStatus()
	fmt.Println("Email before update", newUser.Email)
	newUser.updateEmail()
	fmt.Println("Email after update", newUser.Email)
}

func (u User) getStatus() {
	fmt.Println(u.Status)
}

func (u *User) updateEmail() {
	u.Email = "shibin@outlook.com"
}
