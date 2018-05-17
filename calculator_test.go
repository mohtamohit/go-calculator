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

}

func TestCanMultiply(t *testing.T) {

}

func TestCanDivide(t *testing.T) {

}
