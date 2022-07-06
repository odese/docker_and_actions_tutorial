package utils

// CompareTwoArrays compares two uint64 arrays, and order does not matter
func CompareTwoArrays(arr1, arr2 []uint64) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for _, e := range arr1 {
		if IsUint64ElementInArray(e, arr2) {
			continue
		} else {
			return false
		}
	}

	return true
}

// IsUint64ElementInArray checks if a uint64 element is in a uint64 array
func IsUint64ElementInArray(e uint64, arr []uint64) bool {
	for _, x := range arr {
		if x == e {
			return true
		}
	}
	return false
}
