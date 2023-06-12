package main

import (
	"fmt"
	"time"
)

var counter int

func main() {
	counter = 0

	go incrementCounter()
	go printCounter()

	// Wait for goroutines to finish
	time.Sleep(2 * time.Second)
}

func incrementCounter() {
	for i := 0; i < 10; i++ {
		// Read the current value of counter
		value := counter

		// Increment the counter by 1
		value++

		// Simulate some processing time
		time.Sleep(time.Millisecond)

		// Write the updated value back to counter
		counter = value
	}
}

func printCounter() {
	for i := 0; i < 10; i++ {
		// Read the current value of counter
		value := counter

		// Print the current value
		fmt.Println(value)

		// Simulate some processing time
		time.Sleep(time.Millisecond)
	}
}
