// Package calculator provides basic arithmetic operations
package main

import (
	"errors"
)

// Add returns the sum of two numbers
func Add(a, b float64) float64 {
	return a + b
}

// Subtract returns the difference between two numbers
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply returns the product of two numbers
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide returns the quotient of two numbers
// Returns an error if the divisor is zero
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// Power returns the base raised to the power of exponent
func Power(base, exponent float64) float64 {
	if exponent == 0 {
		return 1
	}
	
	result := base
	for i := 1; i < int(exponent); i++ {
		result *= base
	}
	
	return result
}

// IsEven checks if a number is even
func IsEven(n int) bool {
	return n%2 == 0
}

// Max returns the larger of two numbers
func Max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// Min returns the smaller of two numbers
func Min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}