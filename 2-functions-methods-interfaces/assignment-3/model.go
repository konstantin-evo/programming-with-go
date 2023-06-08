package main

import (
	"fmt"
)

// Animal struct to hold animal information
type Animal struct {
	food       string
	locomotion string
	noise      string
}

// Eat method to print the animal's food
func (a Animal) Eat() {
	fmt.Println(a.food)
}

// Move method to print the animal's locomotion
func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

// Speak method to print the animal's spoken sound
func (a Animal) Speak() {
	fmt.Println(a.noise)
}
