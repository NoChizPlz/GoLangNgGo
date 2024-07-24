package main

import "fmt"

func main() {
	var a, b, c, d int = 1, 2, 3, 4
	var e, f = 5, "John"
	g, h := 7, "Jane"

	var (
		i int
		j int    = 1
		k string = "Doe"
	)

	i = 2

	fmt.Println(a, b, c, d)
	fmt.Println(e, f)
	fmt.Println(g, h)
	println(i, j, k)
}
