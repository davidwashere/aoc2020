package day12

import (
	"fmt"
	"testing"
)

func TestP1(t *testing.T) {
	got := part1("sample.txt")
	fmt.Printf("Got: %v\n", got)
	// want := 37
	// if got != want {
	// 	t.Errorf("Got %v but want %v", got, want)
	// }
}

func TestP1_Actual(t *testing.T) {
	got := part1("input.txt")
	fmt.Printf("Got: %v\n", got)
	// want := 2368
	// if got != want {
	// 	t.Errorf("Got %v but want %v", got, want)
	// }
}

func TestP2(t *testing.T) {
	var got int
	got = part2("sample.txt")
	fmt.Printf("Got: %v\n", got)
	// var want int
	// want = 26
	// if got != want {
	// 	t.Errorf("Got %v but want %v", got, want)
	// }
}

func TestP2_Actual(t *testing.T) {
	got := part2("input.txt")
	fmt.Printf("Got: %v\n", got)
	// want := 2124
	// if got != want {
	// 	t.Errorf("Got %v but want %v", got, want)
	// }
}
