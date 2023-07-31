package main

import "fmt"

func main() {

	Players :=
		map[string]map[string]string{

			"Paweł": map[string]string{
				"Surname":  "Nowak",
				"Position": "Left",
			},
			"Michał": map[string]string{
				"Surname":  "Kowalski",
				"Position": "Right",
			},
		}

	fmt.Println(Players["Paweł"]["Surname"])
}
