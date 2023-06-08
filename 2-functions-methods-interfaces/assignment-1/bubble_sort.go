package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter up to 10 integers (separated by spaces):")

	var numbers []int
	var err error

	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// parse and sort input data
		numbers, err = parseInput(input)
		if err == nil {
			break // exit the loop if input is valid
		}

		fmt.Printf("Invalid input: %s\nPlease try again: ", err)
	}

	bubbleSort(numbers)
	fmt.Println("Array after sorting:", numbers)
}

// BubbleSort performs bubble sort on the given slice of integers
func bubbleSort(numbers []int) {
	n := len(numbers)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				swap(numbers, j)
			}
		}
	}
}

// Swap swaps two elements in the slice using parallel assignment
func swap(numbers []int, i int) {
	numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
}

func parseInput(input string) ([]int, error) {
	numStrings := strings.Fields(input)
	numbers := make([]int, 0, len(numStrings))

	for _, numStr := range numStrings {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, fmt.Errorf("invalid input: %s is not a number", numStr)
		}
		numbers = append(numbers, num)
	}

	if len(numbers) > 10 {
		// If more than 10 numbers were entered, truncate the slice to contain only the first 10 numbers
		numbers = numbers[:10]
	}

	return numbers, nil
}
