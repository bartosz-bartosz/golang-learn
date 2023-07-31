package main

import "fmt"

func main() {
	// Create array
	var SomeArray [5]int

	SomeArray[0] = 9
	SomeArray[1] = 9
	SomeArray[2] = 9
	SomeArray[3] = 2
	SomeArray[4] = 9

	// Create array in one line
	OtherArray := [5]int{0, 2, 4, 6, 8}

	fmt.Println(SomeArray[3])
	fmt.Println(OtherArray[3])

	// Iterate values only
	for _, item := range OtherArray {
		fmt.Println(item + 1)
	}

	fmt.Println("------")

	// Iterate with index
	for index, value := range OtherArray {
		fmt.Println(index, value)
	}

	fmt.Println("------")

	// Slice:

	fmt.Println(OtherArray[1:4][0])
}
