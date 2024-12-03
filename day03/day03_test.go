package day03

import "testing"

func TestFirstHalf(t *testing.T) {
	expected := 161289189
	actual := FirstHalf()

	if expected != actual {
		t.Errorf("TestFirstHalf() = %v; want %v", actual, expected)
	}
}

func TestSecondHalf(t *testing.T) {
	expected := 83595109
	actual := SecondHalf()

	if expected != actual {
		t.Errorf("TestSecondHalf() = %v; want %v", actual, expected)
	}
}
