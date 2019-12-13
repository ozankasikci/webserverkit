package slices

import (
	"math/rand"
	"time"
)

func SliceContainsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func CopyStringSlice(slice []string) []string {
	return append([]string(nil), slice...)
}

func RemoveStringFromSlice(s []string, r string) []string {
	found := false
	j := 0

	copiedSlice := CopyStringSlice(s)
	for i, v := range copiedSlice {
		if v == r {
			found = true
			j = i
		}
	}

	if found {
		return append(copiedSlice[:j], copiedSlice[j+1:]...)
	}
	return s
}

// min: inclusive, max: exclusive
func MakeRange(min, max int) []int {
	a := make([]int, max-min)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func MakeRangeShuffled(min, max int) []int {
	rand.Seed(time.Now().UnixNano())
	a := MakeRange(min, max)
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	return a
}