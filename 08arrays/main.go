package main

import "fmt"

func main() {
	//array declaration,size should be specified
	var fruitArray [4]string
	fruitArray[0] = "apple"
	fruitArray[2] = "apple"
	fruitArray[3] = "apple"

	fmt.Println("FruitArray is", fruitArray)

	//array initialization
	var vegArray = [10]string{"tomato", "beans"}
	fmt.Println("VegArray is", vegArray)
	fmt.Println("Length of VegArray is", len(vegArray))

	// fruitArray[5]="apple" - throws error index 5 out of bounds
}
