package calculator

import "testing"

func logAndAssert(expected, actual int, message string, t *testing.T) {
	if expected != actual {
		t.Log(message, actual)
		t.Fail()
	}
}
func TestCanAdd(t *testing.T) {
	result := Add(2, 3)
	logAndAssert(5, result, "Addition of 2 and 3 should yield 5 but yielded", t)
}

func TestCanSubtract(t *testing.T) {
	result := Subtract(5, 1)
	logAndAssert(4, result, "Subtraction of 1 from 5 should yield 4 but yielded", t)
}

func TestCanMultiply(t *testing.T) {
	result := Multiply(2, 3)
	logAndAssert(6, result, "Multiplication of 2 and 3 should yield 3 but yielded", t)
}

func TestCanDivide(t *testing.T) {
	result := Divide(6, 3)
	logAndAssert(2, result, "Division of 6 by 3 should yield 2 but yielded,", t)
}
