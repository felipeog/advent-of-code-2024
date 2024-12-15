package day15

import "testing"

func TestFirstHalf(t *testing.T) {
	expected := 1457740
	actual := FirstHalf()

	if expected != actual {
		t.Errorf("TestFirstHalf() = %v; want %v", actual, expected)
	}
}

// TODO:
func TestSecondHalf(t *testing.T) {
	expected := -1
	actual := SecondHalf()

	if expected != actual {
		t.Errorf("TestSecondHalf() = %v; want %v", actual, expected)
	}
}
