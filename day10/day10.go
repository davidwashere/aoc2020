package day10

import (
	"aoc2020/util"
	"fmt"
	"sort"
)

func part1(inputfile string) int {
	data, err := util.ReadFileToIntSlice(inputfile)
	// data, err := util.ReadFileToStringSlice(inputfile)
	util.Check(err)

	sort.Ints(data)
	hop1 := 0
	hop3 := 1
	cur := 0
	for _, line := range data {
		diff := line - cur
		// fmt.Printf("cur %v, line %v, diff %v\n", cur, line, diff)
		cur = line
		if diff > 3 {
			fmt.Println("Nops", cur, line)
			break
		}

		if diff == 1 {
			hop1++
			continue
		}

		if diff == 3 {
			hop3++
			continue
		}

	}
	// fmt.Printf("hop1 %v, hop3 %v\n", hop1, hop3)

	return hop1 * hop3
}

func Part2(inputfile string) int {
	data, err := util.ReadFileToIntSlice(inputfile)
	util.Check(err)

	sort.Ints(data)

	r := util.NewRelationator()
	cur := 0

	data = append([]int{0}, data...)
	r.Set(fmt.Sprintf("%v", data[0]), nil)

	for i := 1; i < len(data); i++ {
		line := data[i]
		diff := line - cur
		cur = line
		r.AddChild(fmt.Sprintf("%v", data[i-1]), fmt.Sprintf("%v", line))

		if diff == 1 {
			if i > 1 {
				prev := data[i-2]
				diff = line - prev

				if diff <= 3 {
					r.AddChild(fmt.Sprintf("%v", prev), fmt.Sprintf("%v", line))
				}

				if i > 2 {
					prev = data[i-3]
					diff = line - prev
					if diff <= 3 {
						r.AddChild(fmt.Sprintf("%v", prev), fmt.Sprintf("%v", line))
					}
				}
			}

			continue
		}

		if diff == 2 {
			if i > 1 {
				prev := data[i-2]
				diff = line - prev
				if diff <= 3 {
					r.AddChild(fmt.Sprintf("%v", prev), fmt.Sprintf("%v", line))
				}
			}
			continue
		}

		if diff == 3 {
			continue
		}
	}
	last := data[len(data)-1]
	lastS := fmt.Sprintf("%v", last)
	r.AddChild(lastS, fmt.Sprintf("%v", last+3))

	// fmt.Println(hop1, hop2, hop3)

	// sum := 1
	// for i := len(data) - 1; i >= 0; i-- {
	// 	curS := fmt.Sprintf("%v", data[i])
	// 	children := r.GetChildren(curS)
	// 	// fmt.Println(children)

	// 	if len(children) > 1 {
	// 		sum += len(children)
	// 	}
	// }

	// fmt.Println(r.GetChildren("0"))
	// fmt.Println(r.StringNode("0"))
	// fmt.Println(r.StringNode("1"))
	// fmt.Println(r.StringNode("2"))
	// fmt.Println(r.StringNode("3"))
	// fmt.Println(r.StringNode("4"))

	// Working answer for samples
	// find := fmt.Sprintf("%v", data[len(data)-1])
	// for _, val := range r.GetAllChildren("0") {
	// 	if val == find {
	// 		count++
	// 	}
	// }

	counts := map[string]int{}
	count := 0
	for _, child := range r.GetChildren("0") {
		count += recurChildren(r, counts, child, lastS)
	}

	return count

	// fmt.Println(r.GetAllChildren("0"))

	// return sum
}

func recurChildren(r util.Relationator, counts map[string]int, id string, last string) int {
	val, ok := counts[id]

	if ok {
		// we already 'cost' of this path
		// This is KEY to prevent lengthy re-computation
		// removing this condition will work for samples but takes too long for final
		return val
	}

	if id == last {
		return 1
	}

	for _, child := range r.GetChildren(id) {
		counts[id] += recurChildren(r, counts, child, last)
	}

	return counts[id]
}
