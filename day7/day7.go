package day6

import (
	"aoc2020/util"
	"strconv"
	"strings"
)

type baggy struct {
	name string
	num  int
}

func parseBagContains(contains string) []baggy {
	bags := []baggy{}

	if strings.HasPrefix(contains, "no") {
		return bags
	}

	canHoldS := strings.Split(contains, ",")
	for _, val := range canHoldS {
		val = val[:strings.Index(val, "bag")] // remove 'bag*' from end
		val = strings.TrimSpace(val)

		valS := strings.SplitN(val, " ", 2)

		num, err := strconv.Atoi(valS[0])
		util.Check(err)

		bag := baggy{valS[1], num}
		bags = append(bags, bag)
	}

	return bags
}

func part1(inputfile string) int {
	data, err := util.ReadFileToStringSlice(inputfile)
	util.Check(err)

	bagsThatHold := map[string][]string{} // a map of bag names to the list of other bags which hold this bag name
	for _, line := range data {
		if strings.HasPrefix(line, "shiny gold") {
			continue
		}

		lineSplit := strings.Split(line, "bags contain")
		bagName := strings.TrimSpace(lineSplit[0])
		contains := strings.TrimSpace(lineSplit[1])

		bags := parseBagContains(contains)

		for _, bag := range bags {
			bagsThatHold[bag.name] = append(bagsThatHold[bag.name], bagName)
		}
	}

	result := 0
	processList := []string{}
	alreadyProcessed := map[string]bool{}
	processList = append(processList, bagsThatHold["shiny gold"]...)
	alreadyProcessed["shiny gold"] = true

	for len(processList) > 0 {
		result++
		bagName := processList[0]
		processList = processList[1:]

		alreadyProcessed[bagName] = true // the bag being processed doesn't need to be reprocessed
		for _, canContainMe := range bagsThatHold[bagName] {
			if _, ok := alreadyProcessed[canContainMe]; !ok {
				if util.IsIn(processList, canContainMe) {
					continue
				}

				processList = append(processList, canContainMe)
			}
		}
	}

	return result
}

func part2(inputfile string) int {
	data, err := util.ReadFileToStringSlice(inputfile)
	util.Check(err)

	bagsIHold := map[string][]baggy{}
	for _, line := range data {
		lineSplit := strings.Split(line, "bags contain")
		bagName := strings.TrimSpace(lineSplit[0])
		contains := strings.TrimSpace(lineSplit[1])

		bags := parseBagContains(contains)
		bagsIHold[bagName] = append(bagsIHold[bagName], bags...)
	}

	result := 0
	processList := []baggy{}
	processList = append(processList, baggy{"shiny gold", 0})

	for len(processList) > 0 {
		curBag := processList[0]
		processList = processList[1:]

		if curBag.name != "shiny gold" {
			result++
		}

		for _, heldBag := range bagsIHold[curBag.name] {
			for i := 0; i < heldBag.num; i++ {
				processList = append(processList, heldBag)
			}

		}
	}

	return result
}
