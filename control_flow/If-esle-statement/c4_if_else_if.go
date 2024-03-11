package main

import "fmt"

func main() {
	var day int
	fmt.Println("EnterDay")
	fmt.Scanln(&day)
	if day == 1 {
		fmt.Println("Monday")
	} else if day == 2 {
		fmt.Println("Tuesday")
	} else if day == 3 {
		fmt.Println("Wednesday")
	} else if day == 4 {
		fmt.Println("Thursday")
	} else if day == 5 {
		fmt.Println("Friday")
	} else if day == 6 {
		fmt.Println("Saturday")
	} else if day == 7 {
		fmt.Println("Sunday")
	} else{
		fmt.Println("Wrongday")
	}
}
