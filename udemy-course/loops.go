package main

import "log"

func main() {

	for i := 0; i <= 10; i++ {
		log.Println(i)
	}

	animals := []string{"dog", "fish", "bear", "duck", "cat"} // A SLICE

	// Iterate through animals with indexes...
	for i, animal := range animals {
		log.Println(i, animal)
	}

	// ...and without indexes
	for _, animal := range animals {
		log.Println(animal)
	}

	// Looping through MAP
	jungle := make(map[int]string)
	jungle[1] = "Lion"
	jungle[2] = "Elephant"
	jungle[3] = "Monkey"
	jungle[4] = "Banana"

	for key, beast := range jungle {
		log.Println(key, beast)
	}

	// Looping through custom struct

	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"John", "Smith", "johnsmith@john.com", 20})
	users = append(users, User{"Mike", "Smith", "mikesmith@mike.com", 33})
	users = append(users, User{"John", "Smith", "johnsmith@john.com", 20})

	for _, l := range users {
		log.Println(l.FirstName)
	}
}
