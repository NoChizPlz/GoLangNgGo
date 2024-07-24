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

func process(x int) (squared int, cubed int) {
	squared = square(x)
	cubed = cube(x)
	return
}

func square(number int) (result int) {
	result = number
	for i := 0; i < 1; i++ {
		result *= result
	}
	return
}

func cube(number int) (result int) {
	result = number
	for i := 0; i < 2; i++ {
		result = result * number
	}
	return
}

func factorial_recursion(x float64) (y float64) {
	if x > 0 {
		y = x * factorial_recursion(x-1)
	} else {
		y = 1
	}
	return
}

func main() {
	fmt.Println(sum(1, 2))

	nameAge("Jessica", 99)
	nameAge("John", 12)
	nameAge("Jake", 31)

	fmt.Println(concatStringInt("Sample", 1))

	squared, cubed := process(5)
	fmt.Println("Squared:", squared, "Cubed:", cubed)

	a, _ := process(100)
	fmt.Println(a)

	fmt.Println(factorial_recursion(4))
}
