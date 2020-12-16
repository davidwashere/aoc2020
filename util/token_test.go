package util

import "testing"

func TestParseTokens(t *testing.T) {
	vf := func(got, want interface{}) {
		if got != want {
			t.Errorf("Got %v want %v", got, want)
		}
	}
	str := "departure location: 32-842 or 854-967"
	tokens := ParseTokens(str)

	vf(len(tokens.Strs), 2)
	vf(tokens.Strs[0], "departure location")
	vf(tokens.Strs[1], "or")

	vf(len(tokens.Ints), 4)
	vf(tokens.Ints[0], 32)
	vf(tokens.Ints[1], 842)
	vf(tokens.Ints[2], 854)
	vf(tokens.Ints[3], 967)
}
