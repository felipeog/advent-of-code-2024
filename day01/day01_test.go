package day01

import "testing"

func TestFirstHalf(t *testing.T) {
	expected := 765748
	actual := FirstHalf()

	if expected != actual {
		t.Errorf("TestFirstHalf() = %v; want %v", actual, expected)
	}
}

func TestSecondHalf(t *testing.T) {
	expected := 27732508
	actual := SecondHalf()

	if expected != actual {
		t.Errorf("TestSecondHalf() = %v; want %v", actual, expected)
	}
}
