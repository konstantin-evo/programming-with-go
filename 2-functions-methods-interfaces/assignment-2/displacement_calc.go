package main

import (
	"fmt"
)

// Returns a function that calculates displacement based on time, acceleration, initial velocity, and initial displacement.
func calculateDisplacement(a, vo, so float64) func(float64) float64 {
	return func(time float64) float64 {
		return 0.5*a*time*time + vo*time + so
	}
}

func main() {
	var a, vo, so, time float64

	// Prompt the user to enter values for acceleration, initial velocity, and initial displacement
	fmt.Println("Enter acceleration:")
	fmt.Scanln(&a)

	fmt.Println("Enter initial velocity:")
	fmt.Scanln(&vo)

	fmt.Println("Enter initial displacement:")
	fmt.Scanln(&so)

	// Generate the displacement function
	fn := calculateDisplacement(a, vo, so)

	// Prompt the user to enter a value for time
	fmt.Println("Enter time:")
	fmt.Scanln(&time)

	// Compute the displacement after the entered time
	displacement := fn(time)

	fmt.Println("Displacement after", time, "seconds:", displacement)
}
