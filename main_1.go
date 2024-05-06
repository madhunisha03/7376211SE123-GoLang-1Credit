package main_1

import (
	"fmt"
)

func main() {
	// Variable Declarations
	var x int = 10
	var y float64 = 5.5
	z := "Hello"
	//Data Types
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	// Control Structures
	if x > 5 {
		fmt.Println("x>5")
	} else if x < 5 {
		fmt.Println("x<5")
	} else {
		fmt.Println("x==5")
	}
	for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}
	// Functions
	result := add(10, 20)
	fmt.Println("Result:", result)
}

// Function to add two numbers
func add(a, b int) int {
	return a + b
}
