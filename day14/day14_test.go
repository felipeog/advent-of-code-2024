package day14

import "testing"

func TestFirstHalf(t *testing.T) {
	expected := 218295000
	actual := FirstHalf()

	if expected != actual {
		t.Errorf("TestFirstHalf() = %v; want %v", actual, expected)
	}
}

func TestSecondHalf(t *testing.T) {
	expected := 6870
	actual := SecondHalf()

	if expected != actual {
		t.Errorf("TestSecondHalf() = %v; want %v", actual, expected)
	}
}
