package main

import "fmt"

func main() {
	var a = 15 + 5
	fmt.Println(a)

	var (
		b = 100 + 50
		c = b + 20
		d = b + c
	)
	fmt.Println(d)

	var e = 10 * 5
	fmt.Println(e)

	var f = 10
	fmt.Println(f)

	f += 5

	fmt.Println(f)

	var g = 5
	var h = 6
	fmt.Println(g > h)

}
