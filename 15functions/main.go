package main

import "fmt"

func main() {

	greeting()

	//pass by value
	fmt.Println(adder(1, 2))

	//pass by reference
	var val1 int = 1
	var val2 int = 2
	fmt.Printf("values brfore swap %v and %v\n", val1, val2)
	swap(&val1, &val2)
	fmt.Printf("values after swap %v and %v\n", val1, val2)

	//passing unknown number of argument
	fmt.Println(proAdder(1, 2, 3, 4, 5))

	//return more than one
	r1, r2 := returnTwo()
	fmt.Println(r1, r2)
}

func greeting() {
	fmt.Println("Hello there!")
}

func adder(a int, b int) int {
	return a + b
}

func proAdder(values ...int) int {
	var tot int = 0
	for _, val := range values {
		tot = tot + val
	}
	return tot
}

func returnTwo() (string, string) {
	return "return1", "return2"
}

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
