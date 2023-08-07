package main

import (
	"log"
	"time"
)

// Declare variable with automatically inferred type
// This variable is introduced outside of any function, so it is globally accessible
// "Package-level variable"
var s = "seven"

var firstName string
var lastName string
var phoneNumber string
var age int
var birthDate time.Time

func main() {

	var s2 = "six"
	s := "eight" // This variable has the same name as the package-level "s" - watch out

	log.Println("s is:", s)
	log.Println("s2 is:", s2)

	saySomething("xxx")

}

func saySomething(s3 string) (string, string) {
	log.Println("s from the saySomething func is:", s)
	return s3, "world"
}
