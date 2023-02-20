package main

import "fmt"

func main() {
	//using make
	var number = make(map[int]string)

	number[1] = "one"
	number[2] = "two"
	number[3] = "three"
	number[4] = "four"

	fmt.Println("Map number is ", number)

	//accessing a element 
	fmt.Println(number[3])

	//accessing a non existing element - return null
	fmt.Println(number[6])

	//deleting from map
	delete(number,2)
	fmt.Println("Map number after deletion is ", number)

	//loop through map
	for key,val := range number {
		fmt.Printf("The value for %v is %v \n",key,val)
	}
}