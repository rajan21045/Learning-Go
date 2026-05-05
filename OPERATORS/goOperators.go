package main

import "fmt"

func main() {
	var x int = 20
	var y int = 10

	// Arithmetic Operators
	fmt.Printf("Operator (+) name 'addition' Description 'Adds two values' Result: %d\n", (x + y))
	fmt.Printf("Operator (-) name 'subtraction' Description 'Subtracts two values' Result: %d\n", (x - y))
	fmt.Printf("Operator (*) name 'multiplication' Description 'Multiplies two values' Result: %d\n", (x * y))
	fmt.Printf("Operator (/) name 'division' Description 'Divides two values' Result: %d\n", (x / y))
	fmt.Printf("Operator (%%) name 'modulus' Description 'Remainder of division' Result: %d\n", (x % y))

	// Relational Operators
	fmt.Printf("Operator (==) name 'equal to' Description 'Checks if equal' Result: %t\n", (x == y))
	fmt.Printf("Operator (!=) name 'not equal' Description 'Checks if not equal' Result: %t\n", (x != y))
	fmt.Printf("Operator (>) name 'greater than' Description 'Checks if x > y' Result: %t\n", (x > y))
	fmt.Printf("Operator (<) name 'less than' Description 'Checks if x < y' Result: %t\n", (x < y))
	fmt.Printf("Operator (>=) name 'greater or equal' Description 'Checks if x >= y' Result: %t\n", (x >= y))
	fmt.Printf("Operator (<=) name 'less or equal' Description 'Checks if x <= y' Result: %t\n", (x <= y))

	// Logical Operators
	a := true
	b := false
	fmt.Printf("Operator (&&) name 'logical AND' Description 'True if both are true' Result: %t\n", (a && b))
	fmt.Printf("Operator (||) name 'logical OR' Description 'True if one is true' Result: %t\n", (a || b))
	fmt.Printf("Operator (!) name 'logical NOT' Description 'Reverses boolean' Result: %t\n", (!a))

	// Assignment Operators
	var z int = x
	z += y
	fmt.Printf("Operator (+=) name 'add assign' Description 'Adds and assigns' Result: %d\n", z)

	z -= y
	fmt.Printf("Operator (-=) name 'subtract assign' Description 'Subtracts and assigns' Result: %d\n", z)

	z *= y
	fmt.Printf("Operator (*=) name 'multiply assign' Description 'Multiplies and assigns' Result: %d\n", z)

	z /= y
	fmt.Printf("Operator (/=) name 'divide assign' Description 'Divides and assigns' Result: %d\n", z)

	// Bitwise Operators
	fmt.Printf("Operator (&) name 'AND' Description 'Bitwise AND' Result: %d\n", (x & y))
	fmt.Printf("Operator (|) name 'OR' Description 'Bitwise OR' Result: %d\n", (x | y))
	fmt.Printf("Operator (^) name 'XOR' Description 'Bitwise XOR' Result: %d\n", (x ^ y))
	fmt.Printf("Operator (<<) name 'left shift' Description 'Shift bits left' Result: %d\n", (x << 1))
	fmt.Printf("Operator (>>) name 'right shift' Description 'Shift bits right' Result: %d\n", (x >> 1))
}