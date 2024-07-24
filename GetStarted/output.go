package main

import "fmt"

func main() {
	var a, b string = "James", "Bond"

	fmt.Print(a, b, "\n")

	var c, d string = "Matthew", "Luke"

	fmt.Println(c, d)

	var book, chapter, verse = "Mark", 13, 1

	fmt.Printf("The reading today is found in the book %v which is a %T and the chapter and verse is %v:%v \n\n", book, book, chapter, verse)

	var pi = 3.14
	var words = "Gusto kong sumabog"

	fmt.Printf("%v\n", pi)
	fmt.Printf("%#v\n", pi)
	fmt.Printf("%v%%\n", pi)
	fmt.Printf("%T\n", pi)

	fmt.Printf("%v\n", words)
	fmt.Printf("%#v\n", words)
	fmt.Printf("%#T\n\n", words)

	var numero = 15

	fmt.Printf("%b\n", numero)     //Base 2
	fmt.Printf("%d\n", numero)     //Base 10
	fmt.Printf("%+d\n", numero)    //Base 10 and always show sign
	fmt.Printf("%o\n", numero)     //Base 8
	fmt.Printf("%O\n", numero)     //Base 8 with leading 0o
	fmt.Printf("%x\n", numero)     //Base 16 Lower Case
	fmt.Printf("%X\n", numero)     //Base 16 Upper Case
	fmt.Printf("%#x\n", numero)    //Base 16 with leading 0x
	fmt.Printf("%4d\n", numero)    //Pad with spaces (width 4, right justified)
	fmt.Printf("%-4d\n", numero)   //Pad with spaces (width 4, left justified)
	fmt.Printf("%04d\n\n", numero) //Pad with zeroes (width 4

	var txt = "Butillog"

	fmt.Printf("%s\n", txt)    //Plain String
	fmt.Printf("%q\n", txt)    //Doubled-qouted string
	fmt.Printf("%8s\n", txt)   //plain string (width 8, right justified)
	fmt.Printf("%-8s\n", txt)  //plain string (width 8, left justified)
	fmt.Printf("%x\n", txt)    //hex dump of byte values
	fmt.Printf("% x\n\n", txt) //hex dump of byte values with spaces

	var tru = true
	var char = false

	fmt.Printf("%t\n", tru)
	fmt.Printf("%t\n\n", char)

	var cokefloat = 123.321

	fmt.Printf("%e\n", cokefloat)    //Scientific Notation
	fmt.Printf("%f\n", cokefloat)    //Decimal Point no exponent
	fmt.Printf("%.2f\n", cokefloat)  //Default width, precision 2
	fmt.Printf("%6.2f\n", cokefloat) //Width 6, precision 2
	fmt.Printf("%g\n", cokefloat)    //Exponent as needed, only necessary digits

}
