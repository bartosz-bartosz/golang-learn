package main

import "log"

func main() {
	var myString string
	var myInt int

	myString = "Hi"
	myInt = 8
	mySecondString := "Another string"

	log.Println(myString, myInt, mySecondString)

	// A MAP
	// var myMap2 map[string]string // THIS IS NOT HOW MAP SHOULD BE DECLARED!!!

	myMap := make(map[string]string) // Proper way to create a map
	myMap["dog"] = "Azor"
	myMap["cat"] = "Zajonc"

	myMap["dog"] = "Burek"

	log.Println(myMap["dog"])
	log.Println(myMap["cat"])

	anotherMap := make(map[string]int)
	anotherMap["One"] = 1
	anotherMap["Two"] = 2

	log.Println(anotherMap["One"], anotherMap["Two"], anotherMap["Three"])

	type User struct {
		FirstName string
		LastName  string
	}

	userMap := make(map[string]User)

	me := User{
		FirstName: "Jan",
		LastName:  "Kowalski",
	}

	userMap["me"] = me

	log.Println(userMap["me"].FirstName)

}
