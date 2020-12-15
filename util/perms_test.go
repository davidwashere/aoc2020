package util

import (
	"fmt"
	"testing"
)

// TODO: Add assertions to these tests

func TestPerms(t *testing.T) {
	sample := "abc"

	CharPerms(sample, func(perm string) {
		fmt.Println(perm)
	})

	CharPermsOfLen(sample, 2, func(perm string) {
		fmt.Println(perm)
	})
}

func TestPermsI(t *testing.T) {
	sample := []string{"a", "b", "c"}

	Perms(sample, func(perm []interface{}) {
		fmt.Println(perm)
	})

	PermsString(sample, func(perm []string) {
		fmt.Println(perm)
	})

	PermsStringOfLen(sample, 2, func(perm []string) {
		fmt.Println(perm)
	})

	type cat struct {
		name string
		spec string
	}
	cats := []cat{
		cat{"big", "ugly"},
		cat{"fat", "slow"},
		cat{"hi", "world"},
	}

	Perms(cats, func(perm []interface{}) {
		fmt.Println(perm)
	})

}

func TestAppend(t *testing.T) {
	item := "c"
	ans := []interface{}{"a"}

	var newans []interface{}
	newans = append(newans, ans[:]...)
	// fmt.Println(copy(newans, ans[:]))
	fmt.Println(newans)
	newans = append(newans, item)
	fmt.Println(newans, ans)
}
