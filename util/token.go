package util

import (
	"regexp"
	"strconv"
	"strings"
)

// Tokens represents string and int tokens found after parsing an input
type Tokens struct {
	Strs []string
	Ints []int
}

// ParseTokens will parse string and int values out from `input`. A string value is
// defined as [a-zA-z ]+ (includes spaces). An int value is defined as [0-9]+
func ParseTokens(input string) Tokens {
	strRe := regexp.MustCompile(`[a-zA-Z ]+`)
	tokens := Tokens{}
	for _, val := range strRe.FindAllString(input, -1) {
		val = strings.TrimSpace(val)
		if val != "" {
			tokens.Strs = append(tokens.Strs, val)
		}

	}

	intRe := regexp.MustCompile(`[0-9]+`)
	for _, valS := range intRe.FindAllString(input, -1) {
		val, _ := strconv.Atoi(valS)

		tokens.Ints = append(tokens.Ints, val)
	}

	return tokens

}
