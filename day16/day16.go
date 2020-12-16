package day16

import (
	"aoc2020/util"
	"fmt"
	"strconv"
	"strings"
)

type field struct {
	name        string
	validValues map[int]struct{}
}

type fileT struct {
	fields        []field
	validValues   map[int]struct{}
	yourTicket    []int
	nearbyTickets [][]int
}

func NewFile() fileT {
	return fileT{
		fields:        []field{},
		validValues:   map[int]struct{}{},
		yourTicket:    []int{},
		nearbyTickets: [][]int{},
	}
}

func NewField() field {
	return field{
		name:        "",
		validValues: map[int]struct{}{},
	}
}

func parsefile(inputfile string) fileT {
	data, _ := util.ReadFileToStringSlice(inputfile)

	file := NewFile()

	onFields := true
	onMyTicket := false
	onNearby := false
	for _, line := range data {
		if onFields {
			if line == "" {
				onFields = false
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
				file.validValues[i] = struct{}{}
				f.validValues[i] = struct{}{}
			}

			minMaxS = strings.Split(valsS[1], "-")
			min, _ = strconv.Atoi(minMaxS[0])
			max, _ = strconv.Atoi(minMaxS[1])

			for i := min; i <= max; i++ {
				file.validValues[i] = struct{}{}
				f.validValues[i] = struct{}{}
			}

			file.fields = append(file.fields, f)
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

				file.yourTicket = append(file.yourTicket, num)
			}
		}

		if onNearby {
			thisTicket := []int{}
			valsS := strings.Split(line, ",")

			for _, val := range valsS {
				num, _ := strconv.Atoi(val)
				thisTicket = append(thisTicket, num)
			}

			file.nearbyTickets = append(file.nearbyTickets, thisTicket)
		}

		if strings.HasPrefix(line, "nearby tickets") {
			onNearby = true
		} else if strings.HasPrefix(line, "your ticket") {
			onMyTicket = true
		}
	}

	return file
}

func part1(inputfile string) int {
	file := parsefile(inputfile)

	result := 0
	for _, nearbyTicket := range file.nearbyTickets {
		for _, num := range nearbyTicket {
			if _, found := file.validValues[num]; !found {
				// Invalid Value
				result += num
			}
		}
	}

	return result
}

func part2(inputfile string) int {
	file := parsefile(inputfile)

	// Add every field to every possibility index (probably better way to do this)
	possibilities := [][]field{}
	for i := 0; i < len(file.fields); i++ {
		cpy := []field{}
		cpy = append(cpy, file.fields...)
		possibilities = append(possibilities, cpy)
	}

	// Remove invalid tickets from nearbyTickets
	validTickets := [][]int{}
	for _, nearbyTicket := range file.nearbyTickets {
		invalid := false
		for _, num := range nearbyTicket {
			if _, found := file.validValues[num]; !found {
				invalid = true
			}
		}

		if !invalid {
			validTickets = append(validTickets, nearbyTicket)
		}
	}
	file.nearbyTickets = validTickets

	// For each nearby ticket update possibilies with only valid fields
	for _, thisTicket := range file.nearbyTickets {
		for i, num := range thisTicket {
			validFields := []field{}
			for _, itm := range possibilities[i] {
				if _, ok := itm.validValues[num]; ok {
					validFields = append(validFields, itm)
				}
			}

			possibilities[i] = validFields
		}
	}

	for i, pos := range possibilities {
		fmt.Printf("[%v]: ", i)
		for _, f := range pos {
			fmt.Printf("%v, ", f.name)
		}
		fmt.Println()
	}

	// In very ugly way, filter the possibilities until there is one per index
	solid := map[string]bool{}
	final := map[int]field{}
	for len(final) < len(file.fields) {
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
		fmt.Printf("[%v]: ", i)
		for _, f := range pos {
			fmt.Printf("%v, ", f.name)
		}
		fmt.Println()
	}

	// Find the departure fields, and multiply values
	result := 1
	for i, num := range file.yourTicket {
		if strings.HasPrefix(possibilities[i][0].name, "departure") {
			result *= num
		}
	}
	return result
}
