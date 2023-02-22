package main

import "fmt"

func main() {
	// Example 1
	dayOfWeek := 3

	switch dayOfWeek {
	case 1:
		fmt.Println("Sunday")

	case 2:
		fmt.Println("Monday")

	case 3:
		fmt.Println("Tuesday")
		//keyword fallthrough allows us to execute the next case block without checking its condition.
		fallthrough
	case 4:
		fmt.Println("Wednesday")

	case 5:
		fmt.Println("Thursday")

	case 6:
		fmt.Println("Friday")

	case 7:
		fmt.Println("Saturday")

	default:
		fmt.Println("Invalid day")
	}

	//Example 2 -  can also use multiple values inside a single case block

	day := "Sunday"

	switch day {
	case "Saturday", "Sunday":
		fmt.Println("Weekend")

	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Weekday")

	default:
		fmt.Println("Invalid day")
	}

}
