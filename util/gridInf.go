package util

import (
	"fmt"
)

type infinityGridMinMax struct {
	min int
	max int
}

// InfinityGrid grid with infinte dimensions in all directions
// 2d, 3d, ... 11d, doesn't matter (memory constrained)
type InfinityGrid struct {
	// All Data Points
	// The first key `string` will represent the dimensions behind 2d
	//   e.g. "1" might mean z = 1, "1,2" might mean z=1, w=2
	// The two int keys are the `x` and `y` values of of the grid for that dimension
	data map[string]map[int]map[int]string

	// Default value
	def string

	// Max Extents for every `grid`
	xMinMax infinityGridMinMax
	yMinMax infinityGridMinMax

	// MinMax of every dimension
	dimMinMax []infinityGridMinMax

	// False until a value is set into the grid
	initialized bool

	// When bounds are locked minX, minY, maxX, maxY are not updated dynamically
	// Setting values outside the bounds are ignored
	// Getting values outside the bounds returns default value
	BoundsLocked bool
}

func newInfinityGridMinMax() infinityGridMinMax {
	return infinityGridMinMax{
		min: MaxInt,
		max: MinInt,
	}
}

// NewInfinityGrid .
func NewInfinityGrid(defaultValue string) *InfinityGrid {
	return &InfinityGrid{
		data:    map[string]map[int]map[int]string{},
		def:     defaultValue,
		xMinMax: newInfinityGridMinMax(),
		yMinMax: newInfinityGridMinMax(),
	}
}

// NewInfinityGridFromFile .
func NewInfinityGridFromFile(filename string, defaultValue string) *InfinityGrid {
	grid := NewInfinityGrid(defaultValue)

	x := 0
	y := 0
	err := ParseFileAsString(filename, func(line string) {
		for _, char := range line {
			c := string(char)
			grid.Set(c, x, y)
			x++
		}
		y++
		x = 0
	})

	Check(err)

	return grid
}

func _createDimKey(dims ...int) string {
	result := ""

	if len(dims) == 0 {
		return ""
	}

	lastNonZeroIndex := -1
	for i := len(dims) - 1; i >= 0; i-- {
		if dims[i] != 0 {
			lastNonZeroIndex = i
			break
		}
	}

	for i := 0; i <= lastNonZeroIndex; i++ {
		result += fmt.Sprintf("%v", dims[i])
		if i != lastNonZeroIndex {
			result += ","
		}
	}

	return result
}

func (g *InfinityGrid) GetDimensions() int {
	return len(g.dimMinMax)
}

func (g *InfinityGrid) AddDimension() int {
	minMax := newInfinityGridMinMax()
	minMax.min = 0
	minMax.max = 1
	g.dimMinMax = append(g.dimMinMax, minMax)

	return len(g.dimMinMax)
}

// Set .
func (g *InfinityGrid) Set(val string, x, y int, dims ...int) {
	g.xMinMax.min = Min(g.xMinMax.min, x)
	g.xMinMax.max = Max(g.xMinMax.max, x)
	g.yMinMax.min = Min(g.yMinMax.min, y)
	g.yMinMax.max = Max(g.yMinMax.max, y)

	for i, dim := range dims {
		if i > len(g.dimMinMax)-1 {
			g.dimMinMax = append(g.dimMinMax, newInfinityGridMinMax())
		}
		g.dimMinMax[i].min = Min(g.dimMinMax[i].min, dim)
		g.dimMinMax[i].max = Max(g.dimMinMax[i].max, dim)
	}

	dimKey := _createDimKey(dims...)
	data := g.data

	if _, ok := data[dimKey]; !ok {
		data[dimKey] = map[int]map[int]string{}
	}

	if _, ok := data[dimKey][x]; !ok {
		data[dimKey][x] = map[int]string{}
	}

	data[dimKey][x][y] = val
	g.initialized = true
}

// Get .
func (g *InfinityGrid) Get(x, y int, dims ...int) string {
	if !g.initialized {
		return g.def
	}

	data := g.data
	dimKey := _createDimKey(dims...)

	if _, ok := data[dimKey]; !ok {
		return g.def
	}

	if _, ok := data[dimKey][x]; !ok {
		return g.def
	}

	if _, ok := data[dimKey][x][y]; !ok {
		return g.def
	}

	return data[dimKey][x][y]
}

// Width .
func (g *InfinityGrid) Width() int {
	if !g.initialized {
		return 0
	}

	return Abs(g.xMinMax.max-g.xMinMax.min) + 1
}

// Height .
func (g *InfinityGrid) Height() int {
	if !g.initialized {
		return 0
	}

	return Abs(g.yMinMax.max-g.yMinMax.min) + 1
}

// VisitAll will visit every grid coordinate with extents based on
// grids current width & height
func (g *InfinityGrid) VisitAll(visitFunc func(val string, x int, y int, dims ...int)) {
	allDims := calcAllDims(g.dimMinMax)

	for _, dims := range allDims {
		for y := g.yMinMax.min; y <= g.yMinMax.max; y++ {
			for x := g.xMinMax.min; x <= g.xMinMax.max; x++ {
				visitFunc(g.Get(x, y, dims...), x, y, dims...)
			}
		}
	}
}

func calcAllDims(dimMinMax []infinityGridMinMax) [][]int {
	allDims := &[][]int{}

	calcAllDimsRecur(dimMinMax, 0, allDims, []int{})

	return *allDims
}

func calcAllDimsRecur(dimMinMax []infinityGridMinMax, index int, allDims *[][]int, curDim []int) {
	if index == len(dimMinMax) {
		*allDims = append(*allDims, curDim)
		return
	}

	for i := dimMinMax[index].min; i <= dimMinMax[index].max; i++ {
		var newDim []int
		newDim = append(newDim, curDim...)
		newDim = append(newDim, i)

		calcAllDimsRecur(dimMinMax, index+1, allDims, newDim)
	}
}

// Grow Will extend the min, max of every grid and dimension by amt
// Useful for expanding the extents when using VisitAll
func (g *InfinityGrid) Grow(amt int) {
	g.xMinMax.min -= amt
	g.xMinMax.max += amt
	g.yMinMax.min -= amt
	g.yMinMax.max += amt

	for i := 0; i < len(g.dimMinMax); i++ {
		g.dimMinMax[i].min -= amt
		g.dimMinMax[i].max += amt
	}
}

func (g *InfinityGrid) GetMinX() int {
	return g.xMinMax.min
}

func (g *InfinityGrid) GetMinY() int {
	return g.yMinMax.min
}

func (g *InfinityGrid) GetMaxX() int {
	return g.xMinMax.max
}

func (g *InfinityGrid) GetMaxY() int {
	return g.yMinMax.max
}
