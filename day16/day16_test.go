package day16

import "testing"

func TestFirstHalf(t *testing.T) {
	expected := 74392
	actual := FirstHalf()

	if expected != actual {
		t.Errorf("TestFirstHalf() = %v; want %v", actual, expected)
	}
}

// TODO:
func TestSecondHalf(t *testing.T) {
	expected := 0
	actual := SecondHalf()

	if expected != actual {
		t.Errorf("TestSecondHalf() = %v; want %v", actual, expected)
	}
}
