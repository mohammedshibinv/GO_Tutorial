package main

import (
	"fmt"
	"sort"
)

func main() {
	//declaration
	var languages = []string{"c++", "python", "javascript"}

	//append
	languages = append(languages, "go", "pascal")
	fmt.Println("Languages are",languages)
	fmt.Printf("Type of languages is %T\n", languages)

	//using make keyword
	count := make ([]int,4)

	count[0] = 200
	count[1] = 150
	count[2] = 300
	count[3] = 100

	count = append(count, 500,400)
	fmt.Println(count)

	//checking whether slices is sorted or not, return bool
	fmt.Println(sort.IntsAreSorted(count))

	//sorting slices
	sort.Ints(count)
	fmt.Println(count)

	//remove a value from slice based on index
	var cars = []string{"bmw","audi","porshe","GM","kia","tata"}
	fmt.Println(cars)
	index := 2

	fmt.Println(append(cars[:index],cars[index+1:]...))

}