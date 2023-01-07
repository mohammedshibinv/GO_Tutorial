/*
Basic full date  "Mon Jan 2 15:04:05 MST 2006"
Basic short date "2006/01/02"
AM/PM            "3PM==3pm==15h"
No fraction      "Mon Jan _2 15:04:05 MST 2006"
0s for fraction  "15:04:05.00000"
9s for fraction  "15:04:05.99999999"
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("02/01/2006"))

	createdDate := time.Date(2020,time.March,12,23,23,0,0,time.UTC)
	fmt.Println(createdDate)
}