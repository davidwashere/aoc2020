package day12

import (
	"aoc2020/util"
	"fmt"
	"strconv"
)

var vec = util.NewNormalizedVector

type instr struct {
	op  string
	mag int
}

func part1(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)

	instructions := []instr{}

	for _, line := range data {
		ins := instr{}

		ins.op = line[:1]
		ins.mag, _ = strconv.Atoi(line[1:])

		instructions = append(instructions, ins)
	}

	curX := 0
	curY := 0
	dir := vec(1, 0)

	for _, ins := range instructions {
		if ins.op == "N" {
			curY += ins.mag
		} else if ins.op == "S" {
			curY -= ins.mag
		} else if ins.op == "E" {
			curX += ins.mag
		} else if ins.op == "W" {
			curX -= ins.mag
		} else if ins.op == "L" {
			iteractions := ins.mag / 90
			for i := 0; i < iteractions; i++ {
				if dir.X == 1 && dir.Y == 0 {
					dir = vec(0, 1)
				} else if dir.X == 0 && dir.Y == 1 {
					dir = vec(-1, 0)
				} else if dir.X == -1 && dir.Y == 0 {
					dir = vec(0, -1)
				} else if dir.X == 0 && dir.Y == -1 {
					dir = vec(1, 0)
				}
			}
		} else if ins.op == "R" {
			iteractions := ins.mag / 90
			for i := 0; i < iteractions; i++ {
				if dir.X == 1 && dir.Y == 0 {
					dir = vec(0, -1)
				} else if dir.X == 0 && dir.Y == -1 {
					dir = vec(-1, 0)
				} else if dir.X == -1 && dir.Y == 0 {
					dir = vec(0, 1)
				} else if dir.X == 0 && dir.Y == 1 {
					dir = vec(1, 0)
				}
			}
		} else if ins.op == "F" {
			for i := 0; i < ins.mag; i++ {
				curX, curY = dir.Apply(curX, curY)
			}
		}
		fmt.Println(ins, curX, curY)
	}

	return util.Abs(curX) + util.Abs(curY)
}

func part2(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)

	instructions := []instr{}

	for _, line := range data {
		ins := instr{}

		ins.op = line[:1]
		ins.mag, _ = strconv.Atoi(line[1:])

		instructions = append(instructions, ins)
	}

	curX := 0
	curY := 0
	dir := vec(10, 1)

	for _, ins := range instructions {
		if ins.op == "N" {
			dir = vec(dir.X, dir.Y+ins.mag)
		} else if ins.op == "S" {
			dir = vec(dir.X, dir.Y-ins.mag)
		} else if ins.op == "E" {
			dir = vec(dir.X+ins.mag, dir.Y)
		} else if ins.op == "W" {
			dir = vec(dir.X-ins.mag, dir.Y)
		} else if ins.op == "L" {
			iteractions := ins.mag / 90
			for i := 0; i < iteractions; i++ {
				ty := dir.X
				tx := dir.Y * -1
				dir = vec(tx, ty)
				// if dir.X == 1 && dir.Y == 0 {
				// 	dir = vec(0, 1)
				// } else if dir.X == 0 && dir.Y == 1 {
				// 	dir = vec(-1, 0)
				// } else if dir.X == -1 && dir.Y == 0 {
				// 	dir = vec(0, -1)
				// } else if dir.X == 0 && dir.Y == -1 {
				// 	dir = vec(1, 0)
				// }
			}
		} else if ins.op == "R" {
			iteractions := ins.mag / 90
			for i := 0; i < iteractions; i++ {
				tx := dir.Y
				ty := dir.X * -1
				dir = vec(tx, ty)
				// 	if dir.X == 1 && dir.Y == 0 {
				// 		dir = vec(0, -1)
				// 	} else if dir.X == 0 && dir.Y == -1 {
				// 		dir = vec(-1, 0)
				// 	} else if dir.X == -1 && dir.Y == 0 {
				// 		dir = vec(0, 1)
				// 	} else if dir.X == 0 && dir.Y == 1 {
				// 		dir = vec(1, 0)
				// 	}
			}
		} else if ins.op == "F" {
			for i := 0; i < ins.mag; i++ {
				curX, curY = dir.Apply(curX, curY)
			}
		}
		fmt.Println(ins, curX, curY)
	}

	return util.Abs(curX) + util.Abs(curY)
}
