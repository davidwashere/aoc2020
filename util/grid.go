package util

import (
	"fmt"
)

// Grid represents an infinitly growable (memory dependeant) grid of both positive
// and negative coords with string values
//
type Grid struct {
	// All the data points
	data map[int]map[int]string

	// Default value
	def string

	// Automatically updated as values are set
	maxX int
	minX int
	maxY int
	minY int

	// False until a value is set into the grid
	initialized bool

	// When bounds are locked minX, minY, maxX, maxY are not updated dynamically
	// Setting values outside the bounds are ignored
	// Getting values outside the bounds returns default value
	BoundsLocked bool
}

// NewGrid .
func NewGrid(defaultValue string) *Grid {
	return &Grid{
		data: map[int]map[int]string{},
		def:  defaultValue,
		maxX: MinInt,
		minX: MaxInt,
		maxY: MinInt,
		minY: MaxInt,
	}
}

// NewGridFromFile loads the file into a grid
// it assumes each character in the file represents a data point
// data on the first line is row 0, data on the second line is row 1, and so on
func NewGridFromFile(filename string, defaultValue string) *Grid {
	grid := NewGrid(defaultValue)

	x := 0
	y := 0
	err := ParseFileAsString(filename, func(line string) {
		for _, char := range line {
			c := string(char)
			grid.Set(x, y, c)
			x++
		}
		y++
		x = 0
	})

	Check(err)

	return grid
}

// Set .
func (g *Grid) Set(x, y int, val string) {
	g.minX = Min(g.minX, x)
	g.maxX = Max(g.maxX, x)
	g.minY = Min(g.minY, y)
	g.maxY = Max(g.maxY, y)

	data := g.data

	// If the X coord not found in map, it doesn't exist
	if _, ok := data[x]; !ok {
		data[x] = map[int]string{}
	}

	// go requires a dereference in parenthesis
	data[x][y] = val
	g.initialized = true
}

// Get .
func (g *Grid) Get(x, y int) string {
	if !g.initialized {
		return g.def
	}

	data := g.data

	if _, ok := data[x]; !ok {
		return g.def
	}

	if _, ok := data[x][y]; !ok {
		return g.def
	}

	return data[x][y]
}

// Height returns the full height of the grid taking into account negative coords
func (g *Grid) Height() int {
	if !g.initialized {
		return 0
	}

	return Abs(g.maxY-g.minY) + 1
}

// Width returns the full width of the grid taking into account negative coords
func (g *Grid) Width() int {
	if !g.initialized {
		return 0
	}

	return Abs(g.maxX-g.minX) + 1
}

// GetRow returns the full row of the given y coord, the returned data will include
// values at negative indexes
func (g *Grid) GetRow(y int) []string {
	row := []string{}
	if !g.initialized {
		return row
	}

	for x := g.minX; x <= g.maxX; x++ {
		row = append(row, g.Get(x, y))
	}

	return row
}

// VisitAll will visit every grid coordinate with extents based on
// grids current width & height
func (g *Grid) VisitAll(visitFunc func(x int, y int, val string)) {
	for y := g.minY; y <= g.maxY; y++ {
		for x := g.minX; x <= g.maxX; x++ {
			visitFunc(x, y, g.Get(x, y))
		}
	}
}

// Dump Prints out text representation of grid, assumes each values is a single character
func (g *Grid) Dump() {
	if !g.initialized {
		fmt.Println("Grid Not Initialized")
	}

	fmt.Printf("  ")
	for x := g.minX; x <= g.maxX; x++ {
		if x == 0 {
			fmt.Printf("0")
			break
		} else {
			fmt.Printf(" ")
		}
	}
	fmt.Println()

	fmt.Printf("  \u250c") // Top left
	for x := g.minX; x <= g.maxX; x++ {
		fmt.Printf("\u2500")
	}
	fmt.Printf("\u2510") // Top right

	for y := g.maxY; y >= g.minY; y-- {
		fmt.Println()
		if y == 0 {
			fmt.Printf("0 \u2502")
		} else {
			fmt.Printf("  \u2502")
		}

		for x := g.minX; x <= g.maxX; x++ {
			val := g.Get(x, y)
			if val == "" {
				val = " "
			}
			fmt.Print(val)
			if x == g.maxX {
				fmt.Print("\u2502")
			}
		}
	}
	fmt.Println()
	fmt.Printf("  \u2514") // Bot left
	for x := g.minX; x <= g.maxX; x++ {
		fmt.Printf("\u2500")
	}
	fmt.Printf("\u2518") // Bot right
	fmt.Println()
}

// Grow expands the grid min's and max's by amt
func (g *Grid) Grow(amt int) {
	g.minX -= amt
	g.minY -= amt
	g.maxX += amt
	g.maxY += amt
}

func (g *Grid) GetMinX() int {
	return g.minX
}

func (g *Grid) GetMinY() int {
	return g.minY
}

func (g *Grid) GetMaxX() int {
	return g.maxX
}

func (g *Grid) GetMaxY() int {
	return g.maxY
}

func (g *Grid) SetExtents(minX, minY, maxX, maxY int) {
	g.minX = minX
	g.minY = minY
	g.maxX = maxX
	g.maxY = maxY
}

// SetBounds will set and lock the bounds of the grid
func (g *Grid) SetBounds(minX, minY, maxX, maxY int) error {
	if minX > maxX || minY > maxY {
		return fmt.Errorf("Invalid bounds, min can be greater then max")
	}

	g.minX = minX
	g.minY = minY
	g.maxX = maxX
	g.maxY = maxY

	g.LockBounds()

	return nil
}

// LockBounds locks the bounds of the grid
func (g *Grid) LockBounds() {
	g.BoundsLocked = true
}

// UnlockBounds unlocks the bounds of the grid
func (g *Grid) UnlockBounds() {
	g.BoundsLocked = false
}
