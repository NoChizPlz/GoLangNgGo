package main

import "fmt"

func main() {
	day := 7

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("You must be abnormal")
	}

	number := 10

	switch number {
	case 1, 3, 5, 7, 9:
		fmt.Println("Odd")
	case 2, 4, 6, 8, 10:
		fmt.Println("Even")
	default:
		fmt.Println("Out of range")

	}
}
