package util

import (
	"bufio"
	"os"
	"strconv"
)

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
