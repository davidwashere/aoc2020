package day16

import (
	"aoc2020/util"
	"fmt"
	"strconv"
	"strings"
)

func part1(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)

	validValues := map[int]bool{}

	result := 0
	onFields := true
	onNearby := false
	for _, line := range data {
		if onFields {
			if line == "" {
				onFields = false
				continue
			}

			lineS := strings.Split(line, ": ")
			valsS := strings.Split(lineS[1], " or ")

			minMaxS := strings.Split(valsS[0], "-")
			min, _ := strconv.Atoi(minMaxS[0])
			max, _ := strconv.Atoi(minMaxS[1])

			for i := min; i <= max; i++ {
				validValues[i] = true
			}

			minMaxS = strings.Split(valsS[1], "-")
			min, _ = strconv.Atoi(minMaxS[0])
			max, _ = strconv.Atoi(minMaxS[1])

			for i := min; i <= max; i++ {
				validValues[i] = true
			}
		} else if !onNearby {
			if strings.HasPrefix(line, "nearby tickets") {
				onNearby = true
			} else {
				continue
			}
		} else {
			valsS := strings.Split(line, ",")

			for _, val := range valsS {
				num, _ := strconv.Atoi(val)
				if _, ok := validValues[num]; !ok {
					result += num
				}
			}
		}

	}

	return result
}

type field struct {
	name        string
	validValues map[int]bool
}

func NewField() field {
	return field{
		name:        "",
		validValues: map[int]bool{},
	}
}

func part2(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)

	fields := []field{}
	validValues := map[int]bool{}
	myTicket := []int{}
	possibilities := [][]field{}

	onFields := true
	onMyTicket := false
	onNearby := false
	for _, line := range data {
		if onFields {
			if line == "" {
				onFields = false

				for i := 0; i < len(fields); i++ {
					cpy := []field{}
					cpy = append(cpy, fields...)
					possibilities = append(possibilities, cpy)
				}

				continue
			}
			f := NewField()

			lineS := strings.Split(line, ": ")
			f.name = lineS[0]

			valsS := strings.Split(lineS[1], " or ")

			minMaxS := strings.Split(valsS[0], "-")
			min, _ := strconv.Atoi(minMaxS[0])
			max, _ := strconv.Atoi(minMaxS[1])

			for i := min; i <= max; i++ {
				validValues[i] = true
				f.validValues[i] = true
			}

			minMaxS = strings.Split(valsS[1], "-")
			min, _ = strconv.Atoi(minMaxS[0])
			max, _ = strconv.Atoi(minMaxS[1])

			for i := min; i <= max; i++ {
				validValues[i] = true
				f.validValues[i] = true
			}

			fields = append(fields, f)
			continue
		}

		if onMyTicket {
			if line == "" {
				onMyTicket = false
				continue
			}
			valsS := strings.Split(line, ",")

			for _, val := range valsS {
				num, _ := strconv.Atoi(val)

				myTicket = append(myTicket, num)
			}
		}

		if onNearby {
			thisTicket := []int{}
			valsS := strings.Split(line, ",")

			invalid := false
			for _, val := range valsS {
				num, _ := strconv.Atoi(val)
				if _, ok := validValues[num]; !ok {
					invalid = true
					break
				}
				thisTicket = append(thisTicket, num)
			}

			if invalid {
				continue
			}

			for i, num := range thisTicket {
				valid := []field{}
				for _, itm := range possibilities[i] {
					if _, ok := itm.validValues[num]; ok {
						valid = append(valid, itm)
					}
				}

				possibilities[i] = valid
			}
		}

		if strings.HasPrefix(line, "nearby tickets") {
			onNearby = true
		} else if strings.HasPrefix(line, "your ticket") {
			onMyTicket = true
		}
	}

	solid := map[string]bool{}
	final := map[int]field{}

	for len(final) < len(fields) {
		for i, pos := range possibilities {
			if len(pos) == 1 {
				solid[pos[0].name] = true
				final[i] = pos[0]
				continue
			}

			remain := []field{}
			for _, f := range pos {
				if _, ok := solid[f.name]; !ok {
					remain = append(remain, f)
				}
			}
			possibilities[i] = remain
		}

	}

	for i, pos := range possibilities {
		fmt.Printf("[%v]: %v\n", i, pos[0].name)
	}

	result := 1
	for i, num := range myTicket {
		if strings.HasPrefix(possibilities[i][0].name, "departure") {
			result *= num
		}
	}

	return result
}
