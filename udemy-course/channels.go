package main

import (
	"log"
	"math/rand"
	"time"
)

func RandomNumber(n int) int {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(n)
	return value
}

func CalculateValue(intChannel chan int) {
	randomNumber := RandomNumber(10)
	intChannel <- randomNumber
}

func main() {

	intChannel := make(chan int)
	defer close(intChannel)

	go CalculateValue(intChannel)

	myNum := <-intChannel
	log.Println(myNum)
}
