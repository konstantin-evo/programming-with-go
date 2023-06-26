package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func mainM() {
	fmt.Println("Enter a series of integers (space-separated):")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	numbers := parseInput(input)

	partitionSize := len(numbers) / 4

	var wg sync.WaitGroup
	wg.Add(4)

	// Create channels with buffer size 1 to receive the sorted subarrays from goroutines
	subArrays := make([][]int, 4)
	channels := make([]chan []int, 4)
	for i := 0; i < 4; i++ {
		channels[i] = make(chan []int, 1)
	}

	// Start 4 goroutines to sort 4 subarrays
	for i := 0; i < 4; i++ {
		go sortSubarray(numbers[i*partitionSize:(i+1)*partitionSize], channels[i], &wg, i+1)
	}

	// Collect sorted subarrays from channels
	// Close the channels after collecting the sorted subarrays
	for i := 0; i < 4; i++ {
		subArrays[i] = <-channels[i]
		close(channels[i])
	}

	// Merge sorted subarrays
	merged := mergeArrays(subArrays)

	// Print sorted array
	fmt.Println("Sorted array:")
	fmt.Println(merged)
}

func parseInput(input string) []int {
	trimmedInput := strings.TrimSpace(input)
	values := strings.Split(trimmedInput, " ")

	numbers := make([]int, 0, len(values))

	for _, value := range values {
		num, err := strconv.Atoi(value)
		if err == nil {
			numbers = append(numbers, num)
		}
	}

	return numbers
}

func sortSubarray(arr []int, channel chan []int, wg *sync.WaitGroup, goroutineID int) {
	defer wg.Done()

	sort.Ints(arr)

	fmt.Printf("Goroutine %d sorted subarray: %v\n", goroutineID, arr)

	channel <- arr
}

func mergeArrays(arrays [][]int) []int {
	result := make([]int, 0)

	for _, arr := range arrays {
		result = append(result, arr...)
	}

	sort.Ints(result)

	return result
}
