package day9

import (
	"aoc2020/util"
	"strconv"
)

func part1(inputfile string, bufferSize int) int {
	data, err := util.ReadFileToStringSlice(inputfile)
	util.Check(err)

	buffer := []int{}
	for i := 0; i < bufferSize; i++ {
		val, _ := strconv.Atoi(data[i])
		buffer = append(buffer, val)
	}

	for i := bufferSize; i < len(data); i++ {
		val, _ := strconv.Atoi(data[i])
		found := false
		for j := 0; j < len(buffer); j++ {
			if found {
				break
			}
			for k := 1; k < len(buffer); k++ {
				left := buffer[j]
				right := buffer[k]
				if left+right == val {
					found = true
					break
				}
			}
		}
		if !found {
			return val
		}
		buffer = append(buffer[1:], val)
	}

	return -1
}

func part2(inputfile string, bufferSize int) int {

	find := part1(inputfile, bufferSize)

	// Parsing file again is a inefficient
	data, err := util.ReadFileToStringSlice(inputfile)
	util.Check(err)

	for i := 0; i < len(data); i++ {
		firstI := i
		first, _ := strconv.Atoi(data[i])
		sum := first
		for j := i + 1; j < len(data); j++ {
			last, _ := strconv.Atoi(data[j])
			lastI := j
			sum += last

			if sum == find {
				min := 99999999999999999
				max := -1
				for k := firstI; k <= lastI; k++ {
					val, _ := strconv.Atoi(data[k])
					if val > max {
						max = val
					}
					if val < min {
						min = val
					}
				}
				return min + max
			}

			if sum > find {
				break
			}

		}
	}
	return 0
}
