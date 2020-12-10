package day10

import (
	"aoc2020/util"
	"sort"
)

func part1(inputfile string) int {
	data, err := util.ReadFileToIntSlice(inputfile)
	util.Check(err)

	sort.Ints(data)

	// Add 0 to start
	data = append([]int{0}, data...) // Add 0 as start

	// Add last 3 jolt hop
	last := data[len(data)-1]
	data = append(data, last+3)

	hops := map[int]int{}
	for i := 1; i < len(data); i++ {
		cur := data[i]
		prev := data[i-1]
		diff := cur - prev
		hops[diff] = hops[diff] + 1
	}

	return hops[1] * hops[3]
}

func part2(inputfile string) int {
	data, err := util.ReadFileToIntSlice(inputfile)
	util.Check(err)

	sort.Ints(data)

	r := util.NewRelationator()

	// prepend 0
	data = append([]int{0}, data...)

	// append highest + 3 (not necessary, will never be an additional path)
	last := data[len(data)-1]
	data = append(data, last+3)

	lastIndex := len(data) - 1
	for i := 0; i <= lastIndex; i++ {
		cur := data[i]

		for j := 1; j <= util.Min(lastIndex-i, 3); j++ {
			next := data[i+j]
			if next-cur <= 3 {
				r.AddChild(cur, next)
			}
		}
	}

	counts := map[int]int{}
	count := 0
	for _, child := range r.GetChildren(0) {
		count += recurChildren(r, counts, child.(int), last)
	}

	return count
}

func recurChildren(r util.Relationator, counts map[int]int, id int, last int) int {
	val, ok := counts[id]

	if ok {
		// we already know num sub paths for this id
		// This is KEY to prevent lengthy re-computation
		// removing this condition will work for samples but takes too long for final
		return val
	}

	if id == last {
		return 1
	}

	for _, child := range r.GetChildren(id) {
		counts[id] += recurChildren(r, counts, child.(int), last)
	}

	return counts[id]
}
