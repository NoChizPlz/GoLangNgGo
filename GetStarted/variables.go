package main

import "fmt"

/*
Difference Between var and :=

var
-Can be used inside and outside of functions
-Variable declaration and value assignment can be done separately

:=
Can only be used inside functions
Variable declaration and value assignment cannot be done separately (must be done in the same line)
*/

func main() {
	var a string
	var b = "John"
	c := 2

	a = "Jane"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
