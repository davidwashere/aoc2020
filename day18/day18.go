package day18

import (
	"aoc2020/util"
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

	sum := 0
	for _, line := range data {

		result := calcLine(line)

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

type GStack []string

func (o *GStack) Pop() string {
	index := len(*o) - 1
	ele := (*o)[index]
	*o = (*o)[:index]
	return ele
}

func (o *GStack) Push(i string) {
	*o = append(*o, i)
}

func (o *GStack) Empty() bool {
	return len(*o) == 0
}

func (o *GStack) Peek() string {
	return (*o)[len(*o)-1]
}

func part2(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)

	sum := 0
	for _, line := range data {
		sum += calcLineP2(line)
	}

	return sum
}

func calcLineP2(line string) int {
	line = strings.ReplaceAll(line, " ", "")
	intRe := regexp.MustCompile("[0-9]+")
	parenRe := regexp.MustCompile("[()]+")
	opRe := regexp.MustCompile("[*+]+")

	var opStack GStack
	var outStack GStack

	for _, cbyte := range line {
		char := string(cbyte)
		if intRe.MatchString(char) {
			// is num
			outStack.Push(char)

		} else if parenRe.MatchString(char) {
			// is paran
			if char == "(" {
				opStack.Push("(")
			} else if char == ")" {
				for opStack.Peek() != "(" {
					outStack.Push(opStack.Pop())
				}
				opStack.Pop()
			}

		} else if opRe.MatchString(char) {
			// is + or -
			if opStack.Empty() {
				opStack.Push(char)
				continue
			}
			// + higher then *
			for !opStack.Empty() && (char == "*" && opStack.Peek() == "+") {
				outStack.Push(opStack.Pop())
			}

			opStack.Push(char)
		}
	}

	for !opStack.Empty() {
		outStack.Push(opStack.Pop())
	}

	// fmt.Println(outStack)

	var numStack Stack
	for i := 0; i < len(outStack); i++ {
		char := string(outStack[i])
		if intRe.MatchString(char) {
			num, _ := strconv.Atoi(char)
			numStack.Push(num)
			continue
		}

		if opRe.MatchString(char) {
			right := numStack.Pop()
			left := numStack.Pop()

			if char == "*" {
				numStack.Push(left * right)
			} else if char == "+" {
				numStack.Push(left + right)
			}
		}
	}

	return numStack.Pop()
}
