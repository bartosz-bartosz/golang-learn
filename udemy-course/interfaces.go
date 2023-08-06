package main

import "log"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Azor",
		Breed: "German Shepherd",
	}

	PrintInfo(&dog) // Pointer

	gorilla := Gorilla{
		Name:          "King",
		Color:         "Black",
		NumberOfTeeth: 38,
	}

	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {
	log.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}

func (d *Dog) Says() string { // Notice the pointer!
	return "Woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (d *Gorilla) Says() string { // Notice the pointer!
	return "Ugh"
}

func (d *Gorilla) NumberOfLegs() int {
	return 2
}
