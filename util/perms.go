package util

import "reflect"

// PermsString same as `Perms` but internally converts data to `string`
func PermsString(inSlice interface{}, f func([]string)) {
	PermsStringOfLen(inSlice, -1, f)
}

// PermsStringOfLen same as `PermsOfLen` but internally converts data to `string`
func PermsStringOfLen(inSlice interface{}, maxlen int, f func([]string)) {
	PermsOfLen(inSlice, maxlen, func(perm []interface{}) {
		conv := make([]string, len(perm))
		for i := 0; i < len(perm); i++ {
			conv[i] = perm[i].(string)
		}
		f(conv)
	})
}

// Perms will call `f` for every permutation of the items in `inSlice`
func Perms(inSlice interface{}, f func([]interface{})) {
	PermsOfLen(inSlice, -1, f)
}

// PermsOfLen same as `Perms` but will limit permutations to the specific length
func PermsOfLen(inSlice interface{}, maxlen int, f func([]interface{})) {
	slice, success := convToType(inSlice, reflect.Slice)

	if !success {
		panic("Slice conversion no worky")
	}

	var convertedData []interface{}
	length := slice.Len()
	convertedData = make([]interface{}, length)
	for i := 0; i < length; i++ {
		convertedData[i] = slice.Index(i).Interface()
	}

	if maxlen > length || maxlen < 0 {
		maxlen = length
	}

	perms(convertedData, nil, maxlen, f)
}

// convToType will return a 'reflected' value of the kind specified
// ref: https://ahmet.im/blog/golang-take-slices-of-any-type-as-input-parameter/
func convToType(arg interface{}, kind reflect.Kind) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)
	if val.Kind() == kind {
		ok = true
	}
	return
}

// perms will recursively find every permutation for `data`, and when a permutation is
// found will call `f`.  Permutation length is limited by `l`
func perms(data []interface{}, ans []interface{}, l int, f func([]interface{})) {
	if ans != nil && len(ans) == l {
		f(ans)
		return
	}

	for i := 0; i < len(data); i++ {
		item := data[i]

		var rod []interface{}
		rod = append(rod, data[0:i]...)
		rod = append(rod, data[i+1:]...)

		var newans []interface{}
		newans = append(newans, ans...)
		newans = append(newans, item)

		perms(rod, newans, l, f)
	}
}

// CharPerms same as `Perms` but looks at characters inside a string (instead of expecting a slice)
func CharPerms(str string, f func(string)) {
	charPerms(str, "", len(str), f)
}

// CharPermsOfLen same as `CharPerms` but will limit permutations to the specific length
func CharPermsOfLen(str string, maxlen int, f func(string)) {
	if maxlen > len(str) {
		maxlen = len(str)
	}
	charPerms(str, "", maxlen, f)
}

// charPerms will recursively find every permutation for `data`, and when a permutation is
// found will call `f`.  Permutation length is limited by `l`
// ref: https://www.geeksforgeeks.org/print-all-the-permutations-of-a-string-without-repetition-using-collections-in-java/
func charPerms(str string, ans string, l int, f func(string)) {
	if len(ans) == l {
		f(ans)
		return
	}

	for i := 0; i < len(str); i++ {
		chr := string(str[i])
		ros := str[0:i] + str[i+1:]
		charPerms(ros, ans+chr, l, f)
	}
}
