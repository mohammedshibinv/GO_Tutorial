package main

import "fmt"

func main() {
	mark := 75
	var result string

	if mark >= 75 {
		result = "Good"
	} else if mark >= 50 {
		result = "Average"
	} else {
		result = "Bad"
	}
	fmt.Println(result)

	if i:=5;i%5==0 {
		fmt.Println("Multiple of 5")
	} else {
		fmt.Println("Not a multiple of 5")
	}
}