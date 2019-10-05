package slices

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
