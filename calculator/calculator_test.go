package calculator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	require.Equal(t, 5, result, "Addition of 2 and 3 should yield 5 but yieled %d", result)
}

func TestSubtract(t *testing.T) {
	result := Subtract(5, 1)
	require.Equal(t, 4, result, "Subtraction of 1 from 5 should yield 4 but yielded %d", result)
}

func TestMultiply(t *testing.T) {
	result := Multiply(2, 3)
	require.Equal(t, 6, result, "Multiplication of 2 and 3 should yield 6 but yielded %d", result)
}

func TestDivide(t *testing.T) {
	result, err := Divide(6, 3)
	if err != nil {
		t.Log("An error occured: ", err.Error())
	}
	require.Equal(t, 2, result, "Division of 6 and 3 should yield 2 but yielded %d", result)
}

func TestDivisionByZeroNotAllowed(t *testing.T) {
	_, err := Divide(6, 0)
	require.EqualError(t, err, "Division by zero is not allowed", "Wrong error message")
}
