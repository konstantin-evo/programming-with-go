package main

import "fmt"

// Animal interface to describe the methods of an animal
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Cow struct representing a cow
type Cow struct{}

// Eat method for Cow
func (c Cow) Eat() {
	fmt.Println("grass")
}

// Move method for Cow
func (c Cow) Move() {
	fmt.Println("walk")
}

// Speak method for Cow
func (c Cow) Speak() {
	fmt.Println("moo")
}

// Bird struct representing a bird
type Bird struct{}

// Eat method for Bird
func (b Bird) Eat() {
	fmt.Println("worms")
}

// Move method for Bird
func (b Bird) Move() {
	fmt.Println("fly")
}

// Speak method for Bird
func (b Bird) Speak() {
	fmt.Println("peep")
}

// Snake struct representing a snake
type Snake struct{}

// Eat method for Snake
func (s Snake) Eat() {
	fmt.Println("mice")
}

// Move method for Snake
func (s Snake) Move() {
	fmt.Println("slither")
}

// Speak method for Snake
func (s Snake) Speak() {
	fmt.Println("hsss")
}

// CreateAnimal creates a new animal of the specified type and returns it as an Animal interface
func CreateAnimal(animalType string) (Animal, error) {
	animalMap := map[string]Animal{
		"cow":   Cow{},
		"bird":  Bird{},
		"snake": Snake{},
	}

	animal, found := animalMap[animalType]
	if !found {
		return nil, fmt.Errorf("invalid animal type")
	}

	return animal, nil
}
