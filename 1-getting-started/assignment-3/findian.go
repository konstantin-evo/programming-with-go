package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Print("Enter a string: ")
	_, err := fmt.Scan(&input)
	if (err != nil) {
		fmt.Println("Invalid input:", err)
		return
	}

	if (strings.Contains(input, "i") && strings.Contains(input, "a") && strings.Contains(input, "n")) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}