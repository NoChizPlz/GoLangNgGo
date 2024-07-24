package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}

	status := [2]string{"ripe", "rotten"}
	fruits := [3]string{"banana", "apple", "tomato"}
	for i := 0; i < len(fruits); i++ {
		for j := 0; j < len(status); j++ {
			fmt.Println(status[j], fruits[i])
		}
	}

	for id, val := range fruits {
		fmt.Printf("%v\t%v\n", id, val)
	}

	for _, val := range fruits {
		fmt.Printf("%v\n", val)
	}

	for id, _ := range fruits {
		fmt.Printf("%v\n", id)
	}

}
