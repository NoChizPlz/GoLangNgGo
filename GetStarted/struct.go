package main

import "fmt"

type Car struct {
	brand    string
	model    int
	capacity int
	color    string
	manual   bool
}

func printcar(car Car) {
	fmt.Println("Brand:", car.brand)
	fmt.Println("Model:", car.model)
	fmt.Println("Capacity:", car.capacity)
	fmt.Println("Is it manual? :", car.manual)
}

func main() {

	var vios Car

	vios.brand = "Toyota"
	vios.model = 2024
	vios.capacity = 5
	vios.manual = false

	fmt.Println("Brand:", vios.brand)
	fmt.Println("Model:", vios.model)
	fmt.Println("Capacity:", vios.capacity)
	fmt.Println("Is it manual? :", vios.manual)

	var mirage Car

	mirage.color = "Red"
	mirage.brand = "Misyubibi"
	mirage.manual = true
	mirage.model = 2019
	mirage.capacity = 10

	printcar(mirage)

}
