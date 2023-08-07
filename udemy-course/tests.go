package main

import (
	"errors"
	"log"
)

func main() {
	result, err := divideNums(80.0, 20.0)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Result of the division is", result)
}

func divideNums(x, y float32) (float32, error) {
	var result float32
	if y == 0 {
		return result, errors.New("Division by 0, you fool")
	}
	result = x / y

	return result, nil
}
