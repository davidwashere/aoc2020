package day8

import (
	"aoc2020/util"
	"strconv"
	"strings"
)

type instruction struct {
	op  string
	arg int
}

type handler func(instruction)

type machine struct {
	Accumulator  int
	Pos          int
	Executed     map[int]struct{}
	Instructions []instruction
	Handlers     map[string]handler
	RepeatHook   func() bool
	HandlerHook  func(instruction) handler
}

func NewMachine(i []instruction) *machine {
	m := &machine{}
	m.Executed = map[int]struct{}{}
	m.Instructions = i

	handlers := map[string]handler{}

	handlers["acc"] = m.acc
	handlers["jmp"] = m.jmp
	handlers["nop"] = m.nop
	m.Handlers = handlers

	// Default Repeat Func
	m.RepeatHook = func() bool { return false }

	// Default Handler Decision
	m.HandlerHook = func(i instruction) handler { return m.Handlers[i.op] }

	return m
}

func (m *machine) Run() {
	for m.Next() {
	}
}

func (m *machine) Next() bool {
	if _, ok := m.Executed[m.Pos]; ok {
		if !m.RepeatHook() {
			return false
		}
	}

	m.Executed[m.Pos] = struct{}{}

	if m.Pos >= len(m.Instructions) {
		return false
	}

	i := m.Instructions[m.Pos]

	handler := m.HandlerHook(i)
	handler(i)

	return true
}

func (m *machine) Reset() {
	m.Executed = map[int]struct{}{}
	m.Pos = 0
	m.Accumulator = 0
}

func (m *machine) acc(i instruction) {
	m.Accumulator += i.arg
	m.Pos++
}

func (m *machine) jmp(i instruction) {
	m.Pos += i.arg
}

func (m *machine) nop(i instruction) {
	m.Pos = m.Pos + 1
}

func parseFile(inputfile string) []instruction {
	data, err := util.ReadFileToStringSlice(inputfile)
	util.Check(err)

	instructions := []instruction{}

	for _, line := range data {
		lineS := strings.Split(line, " ")
		i := instruction{}
		i.op = strings.TrimSpace(lineS[0])
		i.arg, _ = strconv.Atoi(lineS[1])
		instructions = append(instructions, i)
	}

	return instructions
}

func part1(inputfile string) int {
	instructions := parseFile(inputfile)

	m := NewMachine(instructions)
	m.Run()

	return m.Accumulator
}

func part2(inputfile string) int {
	instructions := parseFile(inputfile)

	swapMe := indexOfNextJmpNop(instructions, -1)

	m := NewMachine(instructions)

	m.RepeatHook = func() bool {
		swapMe = indexOfNextJmpNop(instructions, swapMe)
		m.Reset()
		return true
	}

	m.HandlerHook = func(i instruction) handler {
		if m.Pos == swapMe {
			if i.op == "jmp" {
				return m.Handlers["nop"]
			}

			if i.op == "nop" {
				return m.Handlers["jmp"]
			}
		}

		return m.Handlers[i.op]
	}

	m.Run()

	return m.Accumulator
}

func indexOfNextJmpNop(ins []instruction, cur int) int {
	for i := cur + 1; i < len(ins); i++ {
		in := ins[i]
		if in.op == "jmp" {
			return i
		}
		if in.op == "nop" && in.arg != 0 {
			return i
		}
	}

	return -1
}
