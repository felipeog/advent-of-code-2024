package day05

import "testing"

func TestFirstHalf(t *testing.T) {
	expected := 4135
	actual := FirstHalf()

	if expected != actual {
		t.Errorf("TestFirstHalf() = %v; want %v", actual, expected)
	}
}

func TestSecondHalf(t *testing.T) {
	expected := 5285
	actual := SecondHalf()

	if expected != actual {
		t.Errorf("TestSecondHalf() = %v; want %v", actual, expected)
	}
}
