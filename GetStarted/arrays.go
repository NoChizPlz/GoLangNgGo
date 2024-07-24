package main

import "fmt"

func main() {
	var arr1 = [3]int{1, 2, 3}
	var arr2 = [5]int{1, 2, 3, 4, 5}

	fmt.Println(arr1, arr2)

	var arr3 = [...]int{1, 2, 3, 4, 5}

	fmt.Println(arr3)

	var arr4 = [4]string{"Matthew", "Mark", "Luke", "John"}

	fmt.Println(arr4)

	fmt.Println(arr4[0], arr3[1], ":", arr2[1])

	prices := [3]int{100, 200, 500}
	prices[1] = 300
	fmt.Println(prices)

	arr5 := [5]int{}              //not initialized
	arr6 := [5]int{1, 2}          //partially initialized
	arr7 := [5]int{1, 2, 3, 4, 5} //fully initialized

	fmt.Println(arr5)
	fmt.Println(arr6)
	fmt.Println(arr7)

	arr8 := [5]int{2: 69, 3: 420}
	arr9 := [...]int{2: 69, 3: 420}

	fmt.Println(arr8)
	fmt.Println(arr9)

	fmt.Println(len(arr8))

}
