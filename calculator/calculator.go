package calculator

import (
	"errors"
)

// Add performs addition of two numbers
func Add(x, y int) int {
	return x + y
}

// Subtract performs subtraction of two numbers
func Subtract(x, y int) int {
	return x - y
}

// Multiply performs multiplication of two numbers
func Multiply(x, y int) int {
	return x * y
}

// Divide performs division of two numbers
func Divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Division by zero is not allowed")
	}
	return x / y, nil
}
