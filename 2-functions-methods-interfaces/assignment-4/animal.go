package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// HandleNewAnimalCommand handles the "newanimal" command by creating a new animal and adding it to the animals map
func HandleNewAnimalCommand(animals map[string]Animal, animalName string, animalType string) {
	animal, err := CreateAnimal(animalType)
	if err != nil {
		fmt.Println(err)
		return
	}

	animals[animalName] = animal
	fmt.Println("Created it!")
}

// HandleQueryCommand handles the "query" command by retrieving the requested information about the specified animal
func HandleQueryCommand(animals map[string]Animal, animalName string, animalAction string) {
	animal, ok := animals[animalName]
	if !ok {
		fmt.Printf("Animal '%s' does not exist.\n", animalName)
		return
	}

	switch animalAction {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Printf("Invalid information requested: %s\n", animalAction)
	}
}

func main() {
	animals := make(map[string]Animal)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")

	for scanner.Scan() {
		input := scanner.Text()
		fields := strings.Fields(input)

		if len(fields) != 3 {
			fmt.Println("Invalid input. Please provide a command with three strings.")
			fmt.Print("> ")
			continue
		}

		command := fields[0]
		animalName := fields[1]
		details := fields[2]

		switch command {
		case "newanimal":
			HandleNewAnimalCommand(animals, animalName, details)
		case "query":
			HandleQueryCommand(animals, animalName, details)
		default:
			fmt.Println("Invalid command.")
		}

		fmt.Print("> ")
	}

	if scanner.Err() != nil {
		fmt.Println("Error reading input:", scanner.Err())
	}
}
