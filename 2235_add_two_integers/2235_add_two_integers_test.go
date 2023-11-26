package _235_add_two_integers

import "testing"

func TestAddTwoIntegers(t *testing.T) {
	result := addTwoIntegers(1, 2)
	expected := 3

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
