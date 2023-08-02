package main

import "log"

type myStruct struct {
	FirstName string
}

// Assign a function to existing struct
// "*myStruct" points to the struct
func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	var myVar myStruct
	myVar.FirstName = "John"

	myVar2 := myStruct{
		FirstName: "Mary",
	}

	log.Println("myVar is set to", myVar.FirstName)          // This prints a value directly
	log.Println("myVar2 is set to", myVar2.printFirstName()) // This uses a struct function to print a value

}
