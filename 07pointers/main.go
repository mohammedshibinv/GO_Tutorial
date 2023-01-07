package main

import "fmt"

func main() {
	//Empty pointer hold nil
	var ptr *int
	fmt.Println("value of pointer ptr", ptr)

	//create pointer using &(reference) , *(value at) used to get value
	var num int = 23
	var newptr = &num

	fmt.Println("Value of pointer is ",newptr)
	fmt.Println("Value of pointer is ",*newptr)

	*newptr = *newptr + 2
	fmt.Println("New value of num is",num)

}
