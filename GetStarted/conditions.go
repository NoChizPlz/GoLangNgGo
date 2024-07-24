package main

import "fmt"

func main() {
	if 20 > 18 {
		fmt.Println("20 is greater than 18, obviously.")
	}

	x := 20
	y := 18

	if x > y {
		fmt.Println("x is greater than y")
	}

	time := 8

	if time <= 12 {
		fmt.Println("Good AM, Sir")
	} else {
		fmt.Println("Good PM, maam")
	}

	time = 13

	if time <= 12 {
		fmt.Println("Good AM, Sir")
	} else {
		fmt.Println("Good PM, maam")
	}

	time = 25

	if time <= 12 {
		fmt.Println("Good AM, Sir")
	} else if time <= 24 {
		fmt.Println("Good PM, maam")
	} else {
		fmt.Println("IDK what time is that")
	}

	time = 19

	if time < 24 {
		if time < 12 {
			fmt.Println("Good AM, Sir")
		} else {
			fmt.Println("Good PM, maam")
		}
	}

}
