package day11

import (
	"aoc2020/util"
)

type change struct {
	x   int
	y   int
	val string
}

type vec = util.Vector

var adjVectors = []vec{
	vec{X: -1, Y: -1, M: 1},
	vec{X: 0, Y: -1, M: 1},
	vec{X: 1, Y: -1, M: 1},
	vec{X: 1, Y: 0, M: 1},
	vec{X: 1, Y: 1, M: 1},
	vec{X: 0, Y: 1, M: 1},
	vec{X: -1, Y: 1, M: 1},
	vec{X: -1, Y: 0, M: 1},
}

func part1(inputfile string) int {
	grid := util.NewGridFromFile(inputfile, "X") // X = out of bounds

	for {
		changes := []change{}

		grid.VisitAll(func(x int, y int, val string) {
			adjOccSeats := adjOccupiedSeats(grid, x, y)
			if val == "L" {
				if adjOccSeats == 0 {
					changes = append(changes, change{x, y, "#"})
				}

			} else if val == "#" {
				if adjOccSeats >= 4 {
					changes = append(changes, change{x, y, "L"})
				}

			}
		})

		if len(changes) == 0 {
			break
		}

		for _, c := range changes {
			grid.Set(c.x, c.y, c.val)
		}
	}

	occSeats := 0
	grid.VisitAll(func(x int, y int, val string) {
		if val == "#" {
			occSeats++
		}
	})

	return occSeats
}

// adjOccupiedSeats counts number of occupied seats immediatly surrounding
// the x, y coord
func adjOccupiedSeats(grid util.Grid, x, y int) int {
	occSeats := 0
	for _, v := range adjVectors {
		if grid.Get(v.Apply(x, y)) == "#" {
			occSeats++
		}
	}

	return occSeats
}

func part2(inputfile string) int {
	grid := util.NewGridFromFile(inputfile, "X")

	for {
		changes := []change{}
		grid.VisitAll(func(x int, y int, val string) {
			adjOctSeats := adjSeenOccupiedSeats(grid, x, y)
			if val == "L" {
				if adjOctSeats == 0 {
					changes = append(changes, change{x, y, "#"})
				}

			} else if val == "#" {
				if adjOctSeats >= 5 {
					changes = append(changes, change{x, y, "L"})
				}

			}
		})

		if len(changes) == 0 {
			break
		}

		for _, c := range changes {
			grid.Set(c.x, c.y, c.val)
		}
	}

	occSeats := 0
	grid.VisitAll(func(x int, y int, val string) {
		if val == "#" {
			occSeats++
		}
	})

	return occSeats
}

// adjSeenOccupiedSeats counts number of occupied seats 'visible' to
// the x, y coord
func adjSeenOccupiedSeats(grid util.Grid, x, y int) int {
	occSeats := 0
	for _, v := range adjVectors {
		curX, curY := v.Apply(x, y)
		for {
			if grid.Get(curX, curY) == "#" {
				occSeats++
				break
			}

			if grid.Get(curX, curY) == "L" || grid.Get(curX, curY) == "X" {
				break
			}

			curX, curY = v.Apply(curX, curY)
		}
	}

	return occSeats
}
