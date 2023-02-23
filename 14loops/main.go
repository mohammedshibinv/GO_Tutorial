package main

import "fmt"

func main() {
	var names = []string{"shibin", "roshan", "lolan"}

	//Loop1
	fmt.Println("Loop type 1")
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	//Loop2
	fmt.Println("Loop type 2")
	for idx := range names {
		fmt.Println(names[idx])
	}

	//Loop3
	fmt.Println("Loop type 3")
	for idx, val := range names {
		fmt.Printf("%v -> %v\n", idx, val)
	}

	//Loop4
	fmt.Println("Loop type 4")
	count := 1
	for count <= 10 {
		fmt.Println(count)
		count++
	}

	//Break Statement
	fmt.Println("Break Statement")
	count = 1
	for count <= 10 {
		if count == 5 {
			break
		}
		fmt.Println(count)
		count++
	}

	//Continue Statement
	fmt.Println("Continue Statement")
	count = 1
	for count <= 10 {
		if count == 5 {
			count++
			continue
		}
		fmt.Println(count)
		count++
	}

	//Labels
	fmt.Println("Labels in GO")
	if count == 10 {
		goto label1
	} else {
		goto label2
	}

label1:
	fmt.Println("Inside label1")
label2:
	fmt.Println("Inside label2")
}
