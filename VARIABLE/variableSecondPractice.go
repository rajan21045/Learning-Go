package main

import "fmt"

// main function - entry point of the program
func main() {

	// Declare variable 'a' with explicit type int
	var a int = 20

	// Short variable declaration (type inferred automatically)
	b := 20

	// Print values of a and b
	fmt.Println(a)
	fmt.Println(b)

	// Multiple variable declarations with inferred types
	numCars := 20        // integer variable
	temperature := 0.0   // float variable

	// Print multiple values in one line
	fmt.Println(numCars, temperature)

	// Different data types
	i := 42              // integer
	f := 3.14            // float
	g := 0.867 + 0.5i    // complex number

	// Print integer, float, and complex values
	fmt.Println(i, f, g)

	// Multiple variable assignment in a single line
	averageOpenRate, displayMessage := 0.23, "is the average open rate of your message."

	// Print both variables
	fmt.Println(averageOpenRate, displayMessage)
}