package main

import "fmt"

func main() {

	// Just print
	fmt.Println("0", "Hello!")

	// Declaring variable - explicit
	var whatToSay string
	whatToSay = "Goodbye 👋"
	fmt.Println("1", whatToSay)

	// Declaring variable - implicit
	somethingToSay := saySomething()
	fmt.Println("2", somethingToSay)

	// Function call
	much, more := sayMore()
	fmt.Println("3", much, more)

}

// Define function
func saySomething() string {
	return "something"
}

func sayMore() (string, string) {
	return "more", "returns"
}
