package main

import "fmt"

func main() {
	fmt.Println("Code begins")

	// Standard variable holding a single string
	var myString string
	myString = "some string"
	fmt.Println(myString)

	// A slice of multiple strings
	var mySlice []string
	mySlice = append(mySlice, "first string")
	fmt.Println(mySlice)

	var mySliceCopy []string
	mySliceCopy = append(mySlice, "second string")
	fmt.Println(mySliceCopy[1])
	fmt.Println(mySlice)

	var sliceOfSlices [][]string
	sliceOfSlices = append(sliceOfSlices, mySlice, mySliceCopy)
	fmt.Println(sliceOfSlices)

	anotherSlice := []string{"one", "two", "three", "four", "five"}
	fmt.Println(anotherSlice[1:3])

}
