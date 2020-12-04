package util

import "fmt"

// Grid representation that supports negative (and positive) indexes
//
// zero is considered a positive value
//
type Grid struct {
	// Top Left  (negative X and positive Y: -x,  y)
	dataTL [][]string

	// Top Right (postiive X and positive Y:  x,  y)
	dataTR [][]string

	// Bot Left  (negative X and negative Y: -x, -y)
	dataBL [][]string

	// Bot Right (postiive X and negative Y:  x, -y)
	dataBR [][]string

	// Default value
	def string

	// Automatically updated as values are set
	maxX int
	minX int
	maxY int
	minY int

	// False until a value is set into the grid
	initialized bool
}

// NewGrid .
func NewGrid(defaultValue string) Grid {
	maxInt := int(^uint(0) >> 1)
	minInt := -maxInt - 1

	return Grid{
		def:  defaultValue,
		maxX: minInt,
		minX: maxInt,
		maxY: minInt,
		minY: maxInt,
	}
}

// NewGridFromFile loads the file into a grid
// it assumes each character in the file represents a data point
// data on the first line is row 0, data on the second line is row 1, and so on
func NewGridFromFile(filename string, defaultValue string) Grid {
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
	data := g.getData(x, y)

	if x < g.minX {
		g.minX = x
	}

	if x > g.maxX {
		g.maxX = x
	}

	if y < g.minY {
		g.minY = y
	}

	if y > g.maxY {
		g.maxY = y
	}

	x = abs(x)
	y = abs(y)

	maxX := len(*data) - 1
	if x > maxX {
		for i := 0; i < (x - maxX); i++ {
			*data = append(*data, []string{})
		}
	}

	maxY := len((*data)[x]) - 1
	if y > maxY {
		for i := 0; i < (y - maxY); i++ {
			(*data)[x] = append((*data)[x], g.def)
		}
	}

	// go requires a dereference in parenthesis
	(*data)[x][y] = val

	g.initialized = true
}

// Get .
func (g *Grid) Get(x, y int) string {
	if !g.initialized {
		return g.def
	}

	data := g.getData(x, y)

	x = abs(x)
	y = abs(y)

	if x > len(*data)-1 {
		return g.def
	}

	if y > len((*data)[x])-1 {
		return g.def
	}

	return (*data)[x][y]
}

// Height returns the full height of the grid taking into account negative coords
func (g *Grid) Height() int {
	if !g.initialized {
		return 0
	}

	return abs(g.minY) + abs(g.maxY) + 1
}

// Width returns the full width of the grid taking into account negative coords
func (g *Grid) Width() int {
	if !g.initialized {
		return 0
	}

	return abs(g.minX) + abs(g.maxX) + 1
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

	for y := g.maxY; y >= g.minY; y-- {
		fmt.Println()
		if y == 0 {
			fmt.Printf("0 ")
		} else {
			fmt.Printf("  ")
		}

		for x := g.minX; x <= g.maxX; x++ {
			val := g.Get(x, y)
			if val == "" {
				val = " "
			}
			fmt.Print(val)
		}
	}
	fmt.Println()
}

func (g *Grid) getData(x, y int) *[][]string {
	posX := true
	if x < 0 {
		posX = false
	}

	posY := true
	if y < 0 {
		posY = false
	}

	// Top Left
	if !posX && posY {
		return &g.dataTL
	}

	// Top Right
	if posX && posY {
		return &g.dataTR
	}

	// Bottom Left
	if !posX && !posY {
		return &g.dataBL
	}

	// Bottom Right
	// if posX && !posY {
	return &g.dataBR
	// }

	// return nil
}

// Abs returns the absolute value of x.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
