package main

import "log"

func main() {

	isTrue := true
	myNum := 80

	var truthString string
	truthString = checkBool(isTrue, myNum)
	log.Println(truthString)

	truthInt := checkSwitch(isTrue)
	log.Println(truthInt)

}

// IF conditions
func checkBool(par bool, par2 int) string {
	if !par || par2 >= 90 {
		return "True!"
	} else {
		return "False :()"
	}
}

// SWITCH statement
func checkSwitch(par bool) int {
	switch par {
	case true:
		log.Println("Condition is true")
		return 1
	case false:
		log.Println("Condition is false")
		return 2

	default:
		return 3
	}
}
