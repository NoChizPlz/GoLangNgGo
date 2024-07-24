package main

import "fmt"

func main() {
	var mapa = map[string]string{"Brand": "Toyota", "Model": "Vios", "Capacity": "5", "Year": "1965"}
	var mapaint = map[string]int{"John": 1, "Jane": 5, "Doe": 89}

	fmt.Printf("mapa\t%v\n", mapa)
	fmt.Printf("mapaint\t%v\n", mapaint)

	var mapamake = make(map[string]int)
	mapamake["John"] = 1
	mapamake["Jane"] = 69

	c := make(map[string]string)

	c["Brand"] = "Misyubibi"
	c["Model"] = "Strada"
	c["Year"] = "2014"
	c["Color"] = ""

	fmt.Printf("mapamake\t%v\n", mapamake)
	fmt.Printf("c\t%v\n", c)

	fmt.Println(c["Brand"])

	c["Model"] = "Mirage"

	fmt.Println(c["Model"])

	delete(c, "Year")

	fmt.Println(c)

	val1, ok1 := c["Brand"]
	val2, ok2 := c["Year"]
	val3, ok3 := c["Color"]
	_, ok4 := c["Model"]

	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)
	fmt.Println(ok4)

	for k, v := range c {
		fmt.Printf("%v : %v ", k, v)
	}

	a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	var b []string
	b = append(b, "one", "two", "three", "four")

	for k, v := range a {
		fmt.Printf("%v : %v", k, v)
	}

	fmt.Println()

	for _, v := range b {
		fmt.Printf("%v : %v", v, a[v])
	}

}
