package util

import (
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

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// IsIn will return true if val is found in slice, false otherwise
func IsIn(slice []string, val string) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}

	return false
}

// AllInMap will return true if all elements of slice are found in map, false otherwise
func AllInMap(slice []string, m map[string]string) bool {
	for _, v := range slice {
		if _, ok := m[v]; !ok {
			return false
		}
	}

	return true
}

// IsHex will return true if the string represents valid hex uint64 (ignores case)
func IsHex(val string) bool {
	_, err := strconv.ParseUint(val, 16, 64)
	if err != nil {
		return false
	}
	return true
}
