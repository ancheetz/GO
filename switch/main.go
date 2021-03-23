package main

import (
	"fmt"
	"time"
)

func main() {

	data := 586
	fmt.Print("Write ", data, " as ")
	switch data {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 4:
		fmt.Println("four")
	case 586:
		fmt.Println("five hundred eighty-six.")
	}

	switch time.Now().Weekday() {
	case time.Monday, time.Wednesday, time.Thursday, time.Tuesday:
		fmt.Println("It's a Weekday")
	default:
		fmt.Println("It's the weekend")
	}

}
