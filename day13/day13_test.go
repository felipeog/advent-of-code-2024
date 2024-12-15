package day13

import "testing"

func TestFirstHalf(t *testing.T) {
	expected := 33921
	actual := FirstHalf()

	if expected != actual {
		t.Errorf("TestFirstHalf() = %v; want %v", actual, expected)
	}
}

func TestSecondHalf(t *testing.T) {
	expected := 82261957837868
	actual := SecondHalf()

	if expected != actual {
		t.Errorf("TestSecondHalf() = %v; want %v", actual, expected)
	}
}