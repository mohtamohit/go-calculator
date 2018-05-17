package calculator

import "testing"

func TestCanAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Log("Addition of 2 and 3 should yield 5 but yielded", result)
		t.Fail()
	}
}

func TestCanSubtract(t *testing.T) {

}

func TestCanMultiply(t *testing.T) {

}

func TestCanDivide(t *testing.T) {

}
