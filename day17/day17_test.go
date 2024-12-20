package day17

import "testing"

func TestFirstHalf(t *testing.T) {
	expected := "1,4,6,1,6,4,3,0,3"
	actual := FirstHalf()

	if expected != actual {
		t.Errorf("TestFirstHalf() = %v; want %v", actual, expected)
	}
}

// TODO:
func TestSecondHalf(t *testing.T) {
	expected := "b"
	actual := SecondHalf()

	if expected != actual {
		t.Errorf("TestSecondHalf() = %v; want %v", actual, expected)
	}
}
