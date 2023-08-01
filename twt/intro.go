package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("2... %20q\n", "banana")
	fmt.Printf("1... %20q\n", "apple")
	fmt.Printf("3... %20q\n", "can")
	fmt.Printf("4... %20q\n", "grapefruit")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err == nil {
		fmt.Println(input)
	}
}
