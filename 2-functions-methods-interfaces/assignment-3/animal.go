package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Define the predefined set of animals
	animals := map[string]Animal{
		"cow":   {food: "grass", locomotion: "walk", noise: "moo"},
		"bird":  {food: "worms", locomotion: "fly", noise: "peep"},
		"snake": {food: "mice", locomotion: "slither", noise: "hsss"},
	}

	// Create a scanner to read user input
	reader := bufio.NewReader(os.Stdin)

	for {
		// Print the prompt
		fmt.Print("> ")

		// Read the user input
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Split the input into animal name and information requested
		inputParts := strings.Fields(input)
		if len(inputParts) != 2 {
			fmt.Println("Invalid input. Please provide an animal name and the information requested.")
			continue
		}

		// Extract the animal name and information requested
		animalName := strings.ToLower(inputParts[0])
		infoRequested := strings.ToLower(inputParts[1])

		// Retrieve the animal from the predefined set
		animal, found := animals[animalName]
		if !found {
			fmt.Println("Animal not found.")
			continue
		}

		// Process the information requested by calling the appropriate method
		switch infoRequested {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("Invalid information requested.")
		}
	}
}
