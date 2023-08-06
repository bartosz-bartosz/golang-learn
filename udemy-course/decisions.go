package main

import "log"

func main() {

	isTrue := false

	var truthString string
	truthString = checkBool(isTrue)
	log.Println(truthString)

}

func checkBool(par bool) string {
	if par {
		return "True!"
	} else {
		return "False :()"
	}
}
