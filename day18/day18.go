package day18

import (
	"aoc2020/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Stack []int

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(i int) {
	*s = append(*s, i)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		panic("Oh noes")
	} else {
		index := len(*s) - 1
		ele := (*s)[index]
		*s = (*s)[:index]
		return ele
	}
}

type OpStack []string

func (s *OpStack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *OpStack) Push(i string) {
	*s = append(*s, i)
}

func (s *OpStack) Pop() string {
	if s.IsEmpty() {
		panic("Oh noes")
	} else {
		index := len(*s) - 1
		ele := (*s)[index]
		*s = (*s)[:index]
		return ele
	}
}

func (s *OpStack) Top() string {
	return (*s)[len(*s)-1]
}

func part1(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)
	// data, _ := util.ReadFileToIntSlice(inputfile)
	// data, _ := util.ReadFileToStringSliceWithDelim(inputfile, "\n\n")
	// grid := util.NewInfinityGridFromFile(inputfile, ".")

	sum := 0
	for _, line := range data {

		result := calcLine(line)

		fmt.Println(result)

		sum += result
	}

	return sum
}

func calcLine(line string) int {
	line = strings.ReplaceAll(line, " ", "")
	intRe := regexp.MustCompile("[0-9]+")
	parenRe := regexp.MustCompile("[()]+")
	opRe := regexp.MustCompile("[*+]+")

	var numStack Stack
	var opStack OpStack

	for _, cbyte := range line {
		char := string(cbyte)
		if intRe.MatchString(char) {
			// is num
			num, _ := strconv.Atoi(char)
			if numStack.IsEmpty() {
				numStack.Push(num)
			} else {
				if opStack.Top() == "(" {
					numStack.Push(num)
				} else {
					op := opStack.Pop()
					lastNum := numStack.Pop()
					if op == "*" {
						numStack.Push(num * lastNum)
					} else if op == "+" {
						numStack.Push(num + lastNum)
					}
				}
			}
		} else if parenRe.MatchString(char) {
			// is paran
			if char == "(" {
				opStack.Push("(")
			} else if char == ")" {
				op := opStack.Pop()
				if !opStack.IsEmpty() && (opStack.Top() == "+" || opStack.Top() == "*") {
					op = opStack.Pop()
					last := numStack.Pop()
					lastLast := numStack.Pop()
					if op == "*" {
						numStack.Push(last * lastLast)
					} else if op == "+" {
						numStack.Push(last + lastLast)
					}
				}
			}
			// paranStack

		} else if opRe.MatchString(char) {
			opStack.Push(char)
			// is + or -
		}
	}

	if len(numStack) > 1 {
		panic("wtf")
	}

	val := numStack.Pop()

	return val
}

func part2(inputfile string) int {
	return 0
}
