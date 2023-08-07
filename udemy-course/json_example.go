package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Fruit struct {
	Name   string `json:"fruit"`
	Size   string `json:"size"`
	Color  string `json:"color"`
	Amount int    `json:"amount"`
}

func main() {
	myJSON := `
	[
		{
			"fruit": "Apple",
			"size": "Large",
			"color": "Red",
			"amount": 10
		},
		{
			"fruit": "Cherry",
			"size": "Small",
			"color": "Red",
			"amount": 40
		}
		
	]`

	var unmarshalled []Fruit
	err := json.Unmarshal([]byte(myJSON), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling JSON", err)
	}

	log.Printf("Unmarshalled: %v", unmarshalled)

	// write json from a struct

	var mySlice []Fruit

	var marshalled Fruit
	marshalled.Name = "Banana"
	marshalled.Color = "Yellow"
	marshalled.Size = "Medium"
	marshalled.Amount = 8

	mySlice = append(mySlice, marshalled)

	var marshalled2 Fruit
	marshalled2.Name = "Raspberry"
	marshalled2.Color = "Pink"
	marshalled2.Size = "Small"
	marshalled2.Amount = 24

	mySlice = append(mySlice, marshalled2)

	newJSON, err := json.MarshalIndent(mySlice, "", "    ")
	if err != nil {
		log.Println("Error marshalling:", err)
	}

	fmt.Println(string(newJSON))

}
