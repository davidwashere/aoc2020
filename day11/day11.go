package day11

import (
	"aoc2020/util"
)

// floor (.)
// empty seat (L)
// occupied seat (#)
type coord struct {
	x int
	y int
}

type change struct {
	x   int
	y   int
	val string
}

var adjVectors = []coord{
	coord{-1, -1},
	coord{0, -1},
	coord{1, -1},
	coord{1, 0},
	coord{1, 1},
	coord{0, 1},
	coord{-1, 1},
	coord{-1, 0},
}

// adjOccupiedSeats counts number of occupied seats immediatly surrounding
// the x, y coord
func adjOccupiedSeats(grid util.Grid, x, y int) int {
	occupiedSeats := 0
	for _, v := range adjVectors {
		if grid.Get(x+v.x, y+v.y) == "#" {
			occupiedSeats++
		}
	}

	return occupiedSeats
}

func part1(inputfile string) int {
	grid := util.NewGridFromFile(inputfile, "X") // X = out of bounds

	for {
		changes := []change{}

		for y := 0; y < grid.Height(); y++ {
			for x := 0; x < grid.Width(); x++ {

				adjOctSeats := adjOccupiedSeats(grid, x, y)
				if grid.Get(x, y) == "L" {
					if adjOctSeats == 0 {
						changes = append(changes, change{x, y, "#"})
					}

				} else if grid.Get(x, y) == "#" {
					if adjOctSeats >= 4 {
						changes = append(changes, change{x, y, "L"})
					}

				}
			}
		}

		if len(changes) == 0 {
			break
		}

		for _, c := range changes {
			grid.Set(c.x, c.y, c.val)
		}
	}

	occupiedSeats := 0
	for y := 0; y < grid.Height(); y++ {
		for x := 0; x < grid.Width(); x++ {
			if grid.Get(x, y) == "#" {
				occupiedSeats++
			}
		}
	}

	return occupiedSeats
}

// adjSeenOccupiedSeats counts number of occupied seats 'visible' to
// the x, y coord
func adjSeenOccupiedSeats(grid util.Grid, x, y int) int {
	occupiedSeats := 0
	for _, c := range adjVectors {
		curX := x + c.x
		curY := y + c.y
		for {
			if grid.Get(curX, curY) == "#" {
				occupiedSeats++
				break
			}

			if grid.Get(curX, curY) == "L" || grid.Get(curX, curY) == "X" {
				break
			}

			curX = curX + c.x
			curY = curY + c.y
		}
	}

	return occupiedSeats
}

func part2(inputfile string) int {
	grid := util.NewGridFromFile(inputfile, "X")

	for {
		changes := []change{}
		for y := 0; y < grid.Height(); y++ {
			for x := 0; x < grid.Width(); x++ {

				adjOctSeats := adjSeenOccupiedSeats(grid, x, y)
				if grid.Get(x, y) == "L" {
					if adjOctSeats == 0 {
						changes = append(changes, change{x, y, "#"})
					}

				} else if grid.Get(x, y) == "#" {
					if adjOctSeats >= 5 {
						changes = append(changes, change{x, y, "L"})
					}
				}
			}
		}

		if len(changes) == 0 {
			break
		}

		for _, c := range changes {
			grid.Set(c.x, c.y, c.val)
		}
	}

	occupiedSeats := 0
	for y := 0; y < grid.Height(); y++ {
		for x := 0; x < grid.Width(); x++ {
			if grid.Get(x, y) == "#" {
				occupiedSeats++
			}
		}
	}

	return occupiedSeats
}
