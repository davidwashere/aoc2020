package day13

import (
	"aoc2020/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type thing struct {
	EarliestTs int
	BusIDs     []int
}

func parseFile(inputfile string) thing {
	data, _ := util.ReadFileToStringSlice(inputfile)

	earlistTs, _ := strconv.Atoi(data[0])
	busIDs := []int{}

	split := strings.Split(data[1], ",")
	for _, it := range split {
		busID, err := strconv.Atoi(it)
		if err != nil {
			continue
		}

		busIDs = append(busIDs, busID)
	}

	return thing{
		EarliestTs: earlistTs,
		BusIDs:     busIDs,
	}
}

func parseFile2(inputfile string) thing {
	data, _ := util.ReadFileToStringSlice(inputfile)

	earlistTs, _ := strconv.Atoi(data[0])
	busIDs := []int{}

	split := strings.Split(data[1], ",")
	for _, it := range split {
		busID, err := strconv.Atoi(it)
		if err != nil {
			busIDs = append(busIDs, -1)
			continue
		}

		busIDs = append(busIDs, busID)
	}

	return thing{
		EarliestTs: earlistTs,
		BusIDs:     busIDs,
	}
}

type bus struct {
	busID  int
	offset int
}

func parseFileBuses(inputfile string) (int, []bus) {
	data, _ := util.ReadFileToStringSlice(inputfile)

	earlistTs, _ := strconv.Atoi(data[0])

	buses := []bus{}

	split := strings.Split(data[1], ",")
	for i, it := range split {
		busID, err := strconv.Atoi(it)
		if err != nil {
			continue
		}

		buses = append(buses, bus{busID, i})
	}

	sort.SliceStable(buses, func(i, j int) bool {
		return buses[i].busID < buses[j].busID
	})

	return earlistTs, buses
}

func part1(inputfile string) int {
	thing := parseFile(inputfile)

	earliestBusID := 0
	cur := thing.EarliestTs
	for earliestBusID == 0 {
		for _, busID := range thing.BusIDs {
			mod := cur % busID
			fmt.Printf("busID %v, time %v, mod %v\n", busID, cur, mod)
			if mod == 0 {
				earliestBusID = busID
				break
			}
		}
		cur++
	}
	cur--
	fmt.Println(cur, earliestBusID)

	minsWait := cur - thing.EarliestTs
	return earliestBusID * minsWait
}

func Part2(inputfile string) int {
	earliest, buses := parseFileBuses(inputfile)

	bigDumpStupidSlowBus := buses[0]

	start := earliest
	for (start+bigDumpStupidSlowBus.offset)%bigDumpStupidSlowBus.busID != 0 {
		start++
	}

	lastBusIndex := len(buses) - 1
	iter := 0
	for {
		iter++

		for i, bus := range buses {
			mod := (start + bus.offset) % bus.busID

			if mod != 0 {
				break
			}

			if mod == 0 && i == lastBusIndex {
				return start
			}
		}

		start = start + bigDumpStupidSlowBus.busID
		if iter%1000000000 == 0 {
			fmt.Println("Bil:", start)
		}
	}
}
