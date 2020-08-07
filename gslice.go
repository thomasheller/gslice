package gslice

func ContainsInt(slice []int, i int) bool {
	for _, element := range slice {
		if element == i {
			return true
		}
	}

	return false
}

func ContainsString(slice []string, s string) bool {
	for _, element := range slice {
		if element == s {
			return true
		}
	}

	return false
}
