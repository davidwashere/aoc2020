package util

import (
	"testing"
)

func TestGetRow(t *testing.T) {
	g := NewGrid(".")
	row := g.GetRow(0)

	if len(row) != 0 {
		t.Errorf("Expected empty slice, but got slice of len: %v", len(row))
	}

	g.Set(0, 0, "A")
	g.Set(1, 0, "B")
	g.Set(2, 0, "C")
	g.Set(0, 1, "D")
	g.Set(1, 1, "E")
	g.Set(2, 1, "F")

	row = g.GetRow(0)
	want := "B"
	got := row[1]

	if want != got {
		t.Errorf("Expected %v but got %v", want, got)
	}

	row = g.GetRow(1)
	want = "F"
	got = row[2]

	if want != got {
		t.Errorf("Expected %v but got %v", want, got)
	}

}

func TestGet(t *testing.T) {
	g := NewGrid(".")

	want := "."
	got := g.Get(889538, 58934)

	if want != got {
		t.Errorf("Expected %v but got %v", want, got)
	}

	g.Set(0, 0, "A")
	g.Set(1, 1, "C")
	g.Set(0, -1, "B")

	want = "."
	got = g.Get(8734, 3984)

	if want != got {
		t.Errorf("Expected %v but got %v", want, got)
	}

	want = "."
	got = g.Get(0, 3984)

	if want != got {
		t.Errorf("Expected %v but got %v", want, got)
	}

	want = "A"
	got = g.Get(0, 0)

	if want != got {
		t.Errorf("Expected %v but got %v", want, got)
	}

	want = "B"
	got = g.Get(0, -1)

	if want != got {
		t.Errorf("Expected %v but got %v", want, got)
	}

}

func TestHeight(t *testing.T) {
	g := NewGrid(".")
	want := 0
	got := g.Height()

	if want != got {
		t.Errorf("Expected %v but got %v", want, got)
	}

	g.Set(0, 0, "A")
	g.Set(1, 1, "C")
	g.Set(0, -1, "B")

	want = 3
	got = g.Height()

	if want != got {
		t.Errorf("Expected %v but got %v", want, got)
	}

	g = NewGrid(".")
	g.Set(0, 2, "A")
	g.Set(0, 3, "B")
	g.Set(0, 4, "C")
	want = 3

	got = g.Height()

	g = NewGrid(".")
	g.Set(0, -2, "A")
	g.Set(0, -3, "B")
	g.Set(0, -4, "C")
	want = 3

	got = g.Height()

	if want != got {
		t.Errorf("Expected %v but got %v", want, got)
	}
}

func TestWidth(t *testing.T) {
	g := NewGrid(".")
	want := 0
	got := g.Width()

	if want != got {
		t.Errorf("Expected %v but got %v", want, got)
	}

	g.Set(0, 0, "A")
	g.Set(1, 1, "C")
	g.Set(-1, -1, "B")
	g.Set(-1, 1, "D")

	want = 3
	got = g.Width()

	if want != got {
		t.Errorf("Expected %v but got %v", want, got)
	}

	g = NewGrid(".")
	g.Set(2, 0, "A")
	g.Set(3, 0, "B")
	g.Set(4, 0, "C")
	want = 3

	got = g.Width()

	if want != got {
		t.Errorf("Expected %v but got %v", want, got)
	}

	g = NewGrid(".")
	g.Set(-2, 0, "A")
	g.Set(-3, 0, "B")
	g.Set(-4, 0, "C")
	want = 3

	got = g.Width()

	if want != got {
		t.Errorf("Expected %v but got %v", want, got)
	}
}

func TestDump(t *testing.T) {
	g := NewGrid(".")
	g.Set(0, 0, "A")
	g.Set(1, 1, "C")
	g.Set(-1, -1, "B")
	g.Set(-1, 1, "D")

	g.Dump()

}
