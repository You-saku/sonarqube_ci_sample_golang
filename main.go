package main

import (
	"fmt"
)

func main() {
	fmt.Println("Go Calculator Demo")
	fmt.Println("==================")
	
	// Basic arithmetic operations
	a, b := 10.0, 3.0
	
	fmt.Printf("Numbers: %.1f and %.1f\n\n", a, b)
	
	// Addition
	sum := Add(a, b)
	fmt.Printf("Addition: %.1f + %.1f = %.1f\n", a, b, sum)
	
	// Subtraction
	difference := Subtract(a, b)
	fmt.Printf("Subtraction: %.1f - %.1f = %.1f\n", a, b, difference)
	
	// Multiplication
	product := Multiply(a, b)
	fmt.Printf("Multiplication: %.1f × %.1f = %.1f\n", a, b, product)
	
	// Division
	quotient, err := Divide(a, b)
	if err != nil {
		fmt.Printf("Division error: %v\n", err)
	} else {
		fmt.Printf("Division: %.1f ÷ %.1f = %.2f\n", a, b, quotient)
	}
	
	// Power
	power := Power(a, 2)
	fmt.Printf("Power: %.1f² = %.1f\n", a, power)
	
	// Min and Max
	min := Min(a, b)
	max := Max(a, b)
	fmt.Printf("Min(%.1f, %.1f) = %.1f\n", a, b, min)
	fmt.Printf("Max(%.1f, %.1f) = %.1f\n", a, b, max)
	
	// Even check
	number := 10
	if IsEven(number) {
		fmt.Printf("%d is even\n", number)
	} else {
		fmt.Printf("%d is odd\n", number)
	}
	
	fmt.Println("\nDivision by zero test:")
	_, err = Divide(10, 0)
	if err != nil {
		fmt.Printf("Error caught: %v\n", err)
	}
}