package day17

import (
	"aoc2020/util"
)

// .
const (
	ACTIVE   = "#"
	INACTIVE = "."
)

type C3D struct {
	x int
	y int
	z int
}

var adj = []C3D{
	// C3D{0, 0, 0}, // init
	C3D{1, 0, 0},
	C3D{1, -1, 0},
	C3D{0, -1, 0},
	C3D{-1, -1, 0},
	C3D{-1, 0, 0},
	C3D{-1, 1, 0},
	C3D{0, 1, 0},
	C3D{1, 1, 0},

	C3D{0, 0, 1},
	C3D{1, 0, 1},
	C3D{1, -1, 1},
	C3D{0, -1, 1},
	C3D{-1, -1, 1},
	C3D{-1, 0, 1},
	C3D{-1, 1, 1},
	C3D{0, 1, 1},
	C3D{1, 1, 1},

	C3D{0, 0, -1},
	C3D{1, 0, -1},
	C3D{1, -1, -1},
	C3D{0, -1, -1},
	C3D{-1, -1, -1},
	C3D{-1, 0, -1},
	C3D{-1, 1, -1},
	C3D{0, 1, -1},
	C3D{1, 1, -1},
}

type chg struct {
	c   C3D
	val string
}

// If a cube is active and exactly 2 or 3 of its neighbors are also active
// the cube remains active. Otherwise, the cube becomes inactive.

// If a cube is inactive but exactly 3 of its neighbors are active
// the cube becomes active. Otherwise, the cube remains inactive.
func part1(inputfile string) int {
	// the index of grid is the Z coord
	grids := map[int]*util.Grid{}
	grids[0] = util.NewGridFromFile(inputfile, INACTIVE)

	// for z, grid := range grids {
	// 	fmt.Printf("[%v] %+v\n", z, grid)
	// }
	// fmt.Println()

	extents := 0
	// each cycle adds a new -1
	for cycle := 0; cycle < 6; cycle++ {
		extents++
		grids[extents] = util.NewGrid(INACTIVE)
		grids[-extents] = util.NewGrid(INACTIVE)

		o := grids[0]
		g := grids[extents]
		g.SetExtents(o.GetMinX(), o.GetMinY(), o.GetMaxX(), o.GetMaxY())
		// fmt.Printf("%+v\n", g)

		g = grids[-extents]
		g.SetExtents(o.GetMinX(), o.GetMinY(), o.GetMaxX(), o.GetMaxY())

		// Expand the grids
		for _, grid := range grids {
			grid.Grow(1)
			// fmt.Printf("[%v] %+v\n", z, grid)
		}

		changes := []chg{}
		for z, grid := range grids {

			grid.VisitAll(func(x, y int, val string) {
				nActive := activeNeightbors(grids, C3D{x, y, z})

				if val == ACTIVE && (nActive < 2 || nActive > 3) {
					changes = append(changes, chg{
						c:   C3D{x, y, z},
						val: INACTIVE,
					})
				} else if val == INACTIVE && nActive == 3 {
					changes = append(changes, chg{
						c:   C3D{x, y, z},
						val: ACTIVE,
					})
				}

			})
		}

		for _, change := range changes {
			set(grids, change)
		}
	}

	result := 0
	for _, grid := range grids {
		// fmt.Printf("z %v, width %v, height %v\n", z, grid.Width(), grid.Height())
		grid.VisitAll(func(x, y int, val string) {
			if val == ACTIVE {
				result++
			}
		})
	}

	return result
}

func get(grids map[int]*util.Grid, c C3D) string {
	if _, ok := grids[c.z]; !ok {
		return INACTIVE
	}

	grid := grids[c.z]

	return grid.Get(c.x, c.y)
}

func set(grids map[int]*util.Grid, c chg) {
	grid := grids[c.c.z]

	grid.Set(c.c.x, c.c.y, c.val)
}

func activeNeightbors(grids map[int]*util.Grid, origin C3D) int {
	count := 0
	for _, v := range adj {
		c := C3D{
			x: origin.x + v.x,
			y: origin.y + v.y,
			z: origin.z + v.z,
		}

		if get(grids, c) == ACTIVE {
			count++
		}
	}

	return count
}

// =============================
type C4D struct {
	x int
	y int
	z int
	w int
}

var adj4d = []C4D{
	// C3D{0, 0, 0}, // init
	C4D{1, 0, 0, 0}, // w = 0
	C4D{1, -1, 0, 0},
	C4D{0, -1, 0, 0},
	C4D{-1, -1, 0, 0},
	C4D{-1, 0, 0, 0},
	C4D{-1, 1, 0, 0},
	C4D{0, 1, 0, 0},
	C4D{1, 1, 0, 0},
	C4D{0, 0, 1, 0}, // z = 1
	C4D{1, 0, 1, 0},
	C4D{1, -1, 1, 0},
	C4D{0, -1, 1, 0},
	C4D{-1, -1, 1, 0},
	C4D{-1, 0, 1, 0},
	C4D{-1, 1, 1, 0},
	C4D{0, 1, 1, 0},
	C4D{1, 1, 1, 0},
	C4D{0, 0, -1, 0}, // Z = -1
	C4D{1, 0, -1, 0},
	C4D{1, -1, -1, 0},
	C4D{0, -1, -1, 0},
	C4D{-1, -1, -1, 0},
	C4D{-1, 0, -1, 0},
	C4D{-1, 1, -1, 0},
	C4D{0, 1, -1, 0},
	C4D{1, 1, -1, 0},
	C4D{0, 0, 0, 1}, // w = 1
	C4D{1, 0, 0, 1},
	C4D{1, -1, 0, 1},
	C4D{0, -1, 0, 1},
	C4D{-1, -1, 0, 1},
	C4D{-1, 0, 0, 1},
	C4D{-1, 1, 0, 1},
	C4D{0, 1, 0, 1},
	C4D{1, 1, 0, 1},
	C4D{0, 0, 1, 1}, // z = 1
	C4D{1, 0, 1, 1},
	C4D{1, -1, 1, 1},
	C4D{0, -1, 1, 1},
	C4D{-1, -1, 1, 1},
	C4D{-1, 0, 1, 1},
	C4D{-1, 1, 1, 1},
	C4D{0, 1, 1, 1},
	C4D{1, 1, 1, 1},
	C4D{0, 0, -1, 1}, // Z = -1
	C4D{1, 0, -1, 1},
	C4D{1, -1, -1, 1},
	C4D{0, -1, -1, 1},
	C4D{-1, -1, -1, 1},
	C4D{-1, 0, -1, 1},
	C4D{-1, 1, -1, 1},
	C4D{0, 1, -1, 1},
	C4D{1, 1, -1, 1},
	C4D{0, 0, 0, -1}, // w = -1
	C4D{1, 0, 0, -1},
	C4D{1, -1, 0, -1},
	C4D{0, -1, 0, -1},
	C4D{-1, -1, 0, -1},
	C4D{-1, 0, 0, -1},
	C4D{-1, 1, 0, -1},
	C4D{0, 1, 0, -1},
	C4D{1, 1, 0, -1},
	C4D{0, 0, 1, -1}, // z = 1
	C4D{1, 0, 1, -1},
	C4D{1, -1, 1, -1},
	C4D{0, -1, 1, -1},
	C4D{-1, -1, 1, -1},
	C4D{-1, 0, 1, -1},
	C4D{-1, 1, 1, -1},
	C4D{0, 1, 1, -1},
	C4D{1, 1, 1, -1},
	C4D{0, 0, -1, -1}, // Z = -1
	C4D{1, 0, -1, -1},
	C4D{1, -1, -1, -1},
	C4D{0, -1, -1, -1},
	C4D{-1, -1, -1, -1},
	C4D{-1, 0, -1, -1},
	C4D{-1, 1, -1, -1},
	C4D{0, 1, -1, -1},
	C4D{1, 1, -1, -1},
}

type chg4d struct {
	c   C4D
	val string
}

func part2(inputfile string) int {
	grids := map[int]map[int]*util.Grid{}
	grids[0] = map[int]*util.Grid{}
	grids[0][0] = util.NewGridFromFile(inputfile, INACTIVE)

	o := grids[0][0]
	extents := 0
	// each cycle adds a new -1
	for cycle := 0; cycle < 6; cycle++ {
		extents++
		for z := -extents; z <= extents; z++ {
			if _, ok := grids[z]; ok {
				grids[z][extents] = util.NewGrid(INACTIVE)
				grids[z][extents].SetExtents(o.GetMinX(), o.GetMinY(), o.GetMaxX(), o.GetMaxY())
				grids[z][-extents] = util.NewGrid(INACTIVE)
				grids[z][-extents].SetExtents(o.GetMinX(), o.GetMinY(), o.GetMaxX(), o.GetMaxY())
			} else { // doesn't exist yet
				grids[z] = map[int]*util.Grid{}
				for w := -extents; w <= extents; w++ {
					grids[z][w] = util.NewGrid(INACTIVE)
					grids[z][w].SetExtents(o.GetMinX(), o.GetMinY(), o.GetMaxX(), o.GetMaxY())
				}
			}
		}

		for w := -extents; w <= extents; w++ {
			for z := -extents; z <= extents; z++ {
				grids[z][w].Grow(1)
			}
		}

		changes := []chg4d{}
		for w := -extents; w <= extents; w++ {
			for z := -extents; z <= extents; z++ {
				grid := grids[z][w]
				grid.VisitAll(func(x, y int, val string) {
					c := C4D{x, y, z, w}
					nActive := activeNeightbors4(grids, c)

					if val == ACTIVE && (nActive < 2 || nActive > 3) {
						changes = append(changes, chg4d{
							c:   c,
							val: INACTIVE,
						})
					} else if val == INACTIVE && nActive == 3 {
						changes = append(changes, chg4d{
							c:   c,
							val: ACTIVE,
						})
					}
				})
			}
		}

		for _, change := range changes {
			set4(grids, change)
		}
	}

	result := 0
	for w := -extents; w <= extents; w++ {
		for z := -extents; z <= extents; z++ {
			grid := grids[z][w]
			grid.VisitAll(func(x, y int, val string) {
				if val == ACTIVE {
					result++
				}
			})
		}
	}

	return result
}

func get4(grids map[int]map[int]*util.Grid, c C4D) string {
	if _, ok := grids[c.z]; !ok {
		return INACTIVE
	}

	if _, ok := grids[c.z][c.w]; !ok {
		return INACTIVE
	}

	grid := grids[c.z][c.w]

	return grid.Get(c.x, c.y)
}

func set4(grids map[int]map[int]*util.Grid, c chg4d) {
	grid := grids[c.c.z][c.c.w]

	grid.Set(c.c.x, c.c.y, c.val)
}

func activeNeightbors4(grids map[int]map[int]*util.Grid, origin C4D) int {
	count := 0
	for _, v := range adj4d {
		c := C4D{
			x: origin.x + v.x,
			y: origin.y + v.y,
			z: origin.z + v.z,
			w: origin.w + v.w,
		}

		if get4(grids, c) == ACTIVE {
			count++
		}
	}

	return count
}
