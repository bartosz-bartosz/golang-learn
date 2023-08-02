package main

import (
	"log"
)

func main() {
	// Just create a string variable
	var someString string
	someString = "Something"
	log.Println("someString is set to: ", someString)

	// Pass a POINTER of the variable instead of it's value
	changeUsingPointer(&someString)
	log.Println("after the pointer change someString is set to: ", someString)
}

// This function takes the POINTER as a parameter (*string).
// Then it sets a new value to the variable in this place in memory
func changeUsingPointer(s *string) {
	log.Println("s is:", s)
	newValue := "Something else"
	*s = newValue
}
