package main

import "fmt"

const PI = 3.14

const Typed int = 1
const Untyped = 2

const (
	A int = 99
	B     = 10.1
	C     = "Math!"
)

func main() {
	fmt.Println(PI)
	fmt.Println(Typed, Untyped)
	fmt.Println(A, B, C)
}
