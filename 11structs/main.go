package main

import "fmt"

type User struct {
	Name   string
	Status bool
	Age    int
}

func main() {
	me := User{"shibin", true, 22}
	fmt.Println("My details : ", me)
	fmt.Printf("My details are %+v \n",me)
	fmt.Println("My name is",me.Name)
}
