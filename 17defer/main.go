package main

import "fmt"

func main() {
	// defer will move the execution of the statement to the very end inside a function.
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")

	fmt.Println("Hello")

	deferFunction()
}

func deferFunction() {
	for i := 1; i < 5; i++ {
		defer fmt.Println(i)
	}
}
