package day17

import (
	"aoc2020/util"
	"fmt"
)

func part1(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)
	// data, _ := util.ReadFileToIntSlice(inputfile)
	// data, _ := util.ReadFileToStringSliceWithDelim(inputfile, "\n\n")

	for _, line := range data {
		tokens := util.ParseTokens(line)
		// ints := util.ParseInts(line)
		// strs := util.ParseStrs(line)
		// words := util.ParseWords(line)

		fmt.Println(tokens)
	}

	result := 0

	return result
}

func part2(inputfile string) int {
	return 0
}
