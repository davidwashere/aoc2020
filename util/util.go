package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// ParseFile parses file line by line calling `lineHandler` with the []byte's of each line
func ParseFile(filename string, lineHandler func([]byte)) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineHandler(scanner.Bytes())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// ParseFileAsString same as `ParseFile` but converts line to a string before calling `lineHandler`
func ParseFileAsString(filename string, lineHandler func(string)) error {
	return ParseFile(filename, func(line []byte) {
		lineHandler(string(line))
	})
}

// ReadFileToStringSlice Read entire file into string slice, one entry per
// line in file
func ReadFileToStringSlice(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

// ReadFileToIntSlice Read entire file into int slice, one entry per
// line in file
func ReadFileToIntSlice(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		item := scanner.Text()
		num, err := strconv.Atoi(item)
		if err != nil {
			return nil, err
		}
		data = append(data, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

// Perm calls f with each permutation of a.
func PermInt(a []int, f func([]int)) {
	pInt(a, f, 0)
}

// Perm the values at index i to len(a)-1.
func pInt(a []int, f func([]int), i int) {
	if i > len(a) {
		f(a)
		return
	}
	pInt(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		pInt(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}
