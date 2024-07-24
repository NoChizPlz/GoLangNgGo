package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func nameAge(name string, age int) {
	fmt.Println("Name:", name, "Age:", age)
}

func concatStringInt(str string, num int) (result string) {
	result = fmt.Sprintf("%s %d", str, num)
	return
}

func main() {
	fmt.Println(sum(1, 2))

	nameAge("Jessica", 99)
	nameAge("John", 12)
	nameAge("Jake", 31)

	fmt.Println(concatStringInt("Sample", 1))
}
