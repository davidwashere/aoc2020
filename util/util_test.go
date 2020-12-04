package util

import "testing"

func TestAbs(t *testing.T) {
	want := 2
	got := Abs(-2)

	if want != got {
		t.Errorf("Expected %v but got %v", want, got)
	}

	want = 9
	got = Abs(9)

	if want != got {
		t.Errorf("Expected %v but got %v", want, got)
	}

	want = 0
	got = Abs(0)

	if want != got {
		t.Errorf("Expected %v but got %v", want, got)
	}
}
