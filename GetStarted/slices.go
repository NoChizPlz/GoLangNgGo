package main

import "fmt"

func main() {
	slices := []int{}

	fmt.Println(len(slices))
	fmt.Println(cap(slices))
	fmt.Println(slices)

	slices2 := []string{"Slices", "Pieces", "Nicest", "Incest"}
	fmt.Println(len(slices2))
	fmt.Println(cap(slices2))
	fmt.Println(slices2)

	array := [5]int{1, 2, 3, 4, 5}
	slices3 := array[2:4]

	fmt.Printf("myslice = %v\n", slices3)
	fmt.Printf("length = %d\n", len(slices3))
	fmt.Printf("capacity = %d\n", cap(slices3))

	slices4 := make([]int, 3, 5)
	fmt.Printf("myslice = %v\n", slices4)
	fmt.Printf("length = %d\n", len(slices4))
	fmt.Printf("capacity = %d\n", cap(slices4))

	slices5 := make([]int, 3)
	fmt.Printf("myslice = %v\n", slices5)
	fmt.Printf("length = %d\n", len(slices5))
	fmt.Printf("capacity = %d\n", cap(slices5))

	slices6 := []int{1, 2, 3, 4, 5}

	fmt.Println(slices6[2])
	fmt.Println(slices6[3])

	slices6[1] = 99

	fmt.Println(slices6[1])

	slices6 = append(slices6, 100, 200)

	fmt.Println(slices6)

	slices6 = append(slices6, slices3...)

	fmt.Println(slices6)

	fmt.Printf("slices6=%v\n", slices6)
	fmt.Printf("length=%d\n", len(slices6))
	fmt.Printf("capacity=%d\n", cap(slices6))

	slices7 := [6]int{1, 2, 3, 4, 5, 6}
	zslice1 := slices7[1:5]

	fmt.Printf("zslice1 = %v\n", zslice1)
	fmt.Printf("length = %d\n", len(zslice1))
	fmt.Printf("capacity = %d\n", cap(zslice1))

	zslice1 = slices7[3:5]

	fmt.Printf("zslice1 = %v\n", zslice1)
	fmt.Printf("length = %d\n", len(zslice1))
	fmt.Printf("capacity = %d\n", cap(zslice1))

	zslice1 = append(zslice1, 99, 98, 100)
	fmt.Printf("zslice1 = %v\n", zslice1)
	fmt.Printf("length = %d\n", len(zslice1))
	fmt.Printf("capacity = %d\n", cap(zslice1))

	alphabhet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q"}
	fmt.Printf("alphabhet = %v\n", alphabhet)
	fmt.Printf("length = %d\n", len(alphabhet))
	fmt.Printf("capacity = %d\n", cap(alphabhet))

	neededAlpha := alphabhet[:len(alphabhet)-10]
	alphaCopy := make([]string, len(neededAlpha))
	copy(alphaCopy, neededAlpha)

	fmt.Printf("alphabhet = %v\n", alphaCopy)
	fmt.Printf("length = %d\n", len(alphaCopy))
	fmt.Printf("capacity = %d\n", cap(alphaCopy))
}
