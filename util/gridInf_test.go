package util

import (
	"testing"
)

func TestInfinityGridThings(t *testing.T) {
	got := _createDimKey(1, 3, 7)
	want := "1,3,7"
	vf(t, got, want)
}

func TestInfinityGridDimKey(t *testing.T) {
	vf(t, _createDimKey(1, 3, 7), "1,3,7")
	vf(t, _createDimKey(), "")
}

func TestInfinityGridSetGet(t *testing.T) {
	grid := NewInfinityGrid(".")
	// 2D
	grid.Set("a", 0, 0)
	grid.Set("b", 1, 1)
	grid.Set("c", -1, -1)

	vf(t, grid.Get(0, 0), "a")
	vf(t, grid.Get(1, 1), "b")
	vf(t, grid.Get(-1, -1), "c")

	// 3D
	grid.Set("x", 0, 0, 1)
	grid.Set("z", -1, -1, -1)
	grid.Set("G", 6, 6, 0)

	vf(t, grid.Get(0, 0, 0), "a")
	vf(t, grid.Get(1, 1, 0), "b")
	vf(t, grid.Get(-1, -1, 0), "c")
	vf(t, grid.Get(0, 0, 1), "x")
	vf(t, grid.Get(-1, -1, -1), "z")
	vf(t, grid.Get(6, 6, 0), "G")
}

func TestInfinityGridWidth(t *testing.T) {
	g := NewInfinityGrid(".")
	vf(t, g.Width(), 0)

	g.Set("A", 0, 0)
	g.Set("C", 1, 1)
	g.Set("B", -1, -1)
	g.Set("D", -1, 1)
	vf(t, g.Width(), 3)

	g = NewInfinityGrid(".")
	g.Set("A", 2, 0)
	g.Set("B", 3, 0)
	g.Set("C", 4, 0)
	vf(t, g.Width(), 3)

	g = NewInfinityGrid(".")
	g.Set("A", -2, 0)
	g.Set("B", -3, 0)
	g.Set("C", -4, 0)
	vf(t, g.Width(), 3)

	g = NewInfinityGrid(".")
	g.Set("A", -2, 0, -1)
	g.Set("B", -3, 0, 1)
	g.Set("C", -4, 0, 20)
	vf(t, g.Width(), 3)
}

func TestInfinityGridHeight(t *testing.T) {
	g := NewInfinityGrid(".")
	vf(t, g.Height(), 0)

	g.Set("A", 0, 0)
	g.Set("C", 1, 1)
	g.Set("B", 0, -1)
	vf(t, g.Height(), 3)

	g = NewInfinityGrid(".")
	g.Set("A", 0, 2)
	g.Set("B", 0, 3)
	g.Set("C", 0, 4)
	vf(t, g.Height(), 3)

	g = NewInfinityGrid(".")
	g.Set("A", 0, -2)
	g.Set("B", 0, -3)
	g.Set("C", 0, -4)
	vf(t, g.Height(), 3)

	g = NewInfinityGrid(".")
	g.Set("A", 0, -2, -1)
	g.Set("B", 0, -3, 1)
	g.Set("C", 0, -4, 20)
	vf(t, g.Height(), 3)
}

func TestAllDims(t *testing.T) {
	dimMinMax := []infinityGridMinMax{
		infinityGridMinMax{-1, 1},
	}

	allDims := calcAllDims(dimMinMax)

	vf(t, len(allDims), 3)
	vf(t, allDims[0][0], -1)
	vf(t, allDims[1][0], 0)
	vf(t, allDims[2][0], 1)

	dimMinMax = []infinityGridMinMax{
		infinityGridMinMax{-1, 1},
		infinityGridMinMax{-1, 1},
	}

	allDims = calcAllDims(dimMinMax)
	vf(t, len(allDims), 9)
	vf(t, allDims[0][0], -1)
	vf(t, allDims[0][1], -1)

	vf(t, allDims[1][0], -1)
	vf(t, allDims[1][1], 0)

	vf(t, allDims[2][0], -1)
	vf(t, allDims[2][1], 1)

	vf(t, allDims[7][0], 1)
	vf(t, allDims[7][1], 0)

	vf(t, allDims[8][0], 1)
	vf(t, allDims[8][1], 1)
}

// func TestVisitAll(t *testing.T) {
// 	g := NewInfinityGrid(".")
// 	g.Set("A", 1, 1, 0)
// 	g.Set("B", -1, -1, 0)
// 	// g.Set("A", 1, 1, 0, 0)
// 	// g.Set("B", -1, -1, 0, 0)
// 	g.Set("C", -1, -1, 1, 1)
// 	g.Set("D", -1, -1, -1, -1)

// 	g.VisitAll(func(val string, x, y int, dims ...int) {
// 		fmt.Printf("%v,%v,%v: %v\n", x, y, dims, val)

// 	})
// }

func TestLockBounds(t *testing.T) {
	g := NewInfinityGrid(".")
	g.Set("A", 0, 0)
	g.LockBounds()
	g.Set("C", 1, 1)
	g.Set("B", -1, -1)

	g.VisitAll2D(func(val string, x, y int) {
		vf(t, val, "A")
	})

	vf(t, g.Height(), 1)
	vf(t, g.Width(), 1)

	g.UnlockBounds()
	g.Set("C", 1, 1)
	g.Set("B", -1, -1)

	vf(t, g.Height(), 3)
	vf(t, g.Width(), 3)
}

func TestSet2DExtents(t *testing.T) {
	g := NewInfinityGrid(".")
	g.Set("A", 0, 0)
	g.Set("B", 3, 3)
	g.SetExtents(-5, -5, 5, 5)
	vf(t, g.Height(), 11)
	vf(t, g.Width(), 11)

	g.SetExtents(0, 0, 2, 2)
	vf(t, g.Get(3, 3), ".")
	g.SetExtents(0, 0, 4, 4)
	vf(t, g.Get(3, 3), "B")
}
