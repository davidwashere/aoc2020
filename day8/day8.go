package day8

import (
	"aoc2020/util"
	"fmt"
	"strconv"
	"strings"
)

type instruction struct {
	op   string
	plus bool
	arg  int
}

func part1(inputfile string) int {
	data, err := util.ReadFileToStringSlice(inputfile)
	util.Check(err)

	instructions := []instruction{}

	for _, line := range data {
		lineS := strings.Split(line, " ")
		op := strings.TrimSpace(lineS[0])

		i := instruction{}
		i.op = op

		argRaw := lineS[1]
		if strings.HasPrefix(argRaw, "+") {
			i.plus = true
		}

		i.arg, _ = strconv.Atoi(argRaw[1:])
		instructions = append(instructions, i)
	}

	executed := map[int]bool{}
	pos := 0
	accumulator := 0

	for {
		i := instructions[pos]

		if _, ok := executed[pos]; ok {
			return accumulator
		}

		executed[pos] = true

		if i.op == "acc" {
			if i.plus {
				accumulator += i.arg
			} else {
				accumulator -= i.arg
			}
			pos++
		} else if i.op == "jmp" {
			if i.plus {
				pos += i.arg
			} else {
				pos -= i.arg
			}
		} else if i.op == "nop" {
			pos++
		}

	}
}

func part2(inputfile string) int {
	data, err := util.ReadFileToStringSlice(inputfile)
	util.Check(err)

	instructions := []instruction{}

	for _, line := range data {
		lineS := strings.Split(line, " ")
		op := strings.TrimSpace(lineS[0])

		i := instruction{}
		i.op = op

		argRaw := lineS[1]
		if strings.HasPrefix(argRaw, "+") {
			i.plus = true
		}

		i.arg, _ = strconv.Atoi(argRaw[1:])
		instructions = append(instructions, i)
	}

	executed := map[int]bool{}
	pos := 0
	accumulator := 0

	for {
		if pos >= len(instructions) {
			break
		}

		i := instructions[pos]

		if _, ok := executed[pos]; ok {
			executed = map[int]bool{}
			changeNext(instructions)
			pos = 0
			accumulator = 0
			continue
		}

		executed[pos] = true

		if i.op == "acc" {
			if i.plus {
				accumulator += i.arg
			} else {
				accumulator -= i.arg
			}
			pos++
		} else if i.op == "jmp" {
			if i.plus {
				pos += i.arg
			} else {
				pos -= i.arg
			}
		} else if i.op == "nop" {
			pos++
		}

	}

	return accumulator
}

var onJump = true
var swapPos = 0

// change the the next jmp to nop, or if already tried that
// chagne the next nop to jmp
func changeNext(ins []instruction) {
	if swapPos > 0 && onJump || !onJump {
		if onJump {
			ins[swapPos].op = "jmp"
		} else {
			ins[swapPos].op = "nop"
		}
		swapPos++
	}

	findOp := "jmp"
	swapTo := "nop"
	if !onJump {
		findOp = "nop"
		swapTo = "jmp"
	}

	for i := swapPos; i < len(ins); i++ {
		in := (ins)[i]
		if in.op == findOp {
			swapPos = i
			(ins)[i].op = swapTo
			return
		}
	}

	fmt.Println("Swapping to nop > jmp")
	onJump = !onJump
	swapPos = 0
}
