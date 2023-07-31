package main

import "fmt"

func main() {
	x, y := 1, 2
	result := add(x, y)

	fmt.Println(result)

}

func add(num1, num2 int) int {
	return num1 + num2
}
