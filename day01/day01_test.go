package day01

import "testing"

func TestPart1SampleInput(t *testing.T) {
	got := part1("sample.txt")
	want := 514579
	t.Logf("Got: %v", got)

	if want != got {
		t.Errorf("Expected %v but got %v", want, got)
	}
}

func TestPart1(t *testing.T) {
	got := part1("input.txt")
	t.Logf("Got: %v", got)
}

func TestPart2SampleInput(t *testing.T) {
	got := part2("sample.txt")
	want := 241861950
	t.Logf("Got: %v", got)

	if want != got {
		t.Errorf("Expected %v but got %v", want, got)
	}
}

func TestPart2(t *testing.T) {
	got := part2("input.txt")
	t.Logf("Got: %v", got)
}
