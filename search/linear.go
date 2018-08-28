package search

// Complexity: O(n)
func LinearSearchUnsorted(data []int, value int) bool {
	for _, item := range data {
		if item == value {
			return true
		}
	}
	return false
}

// Complexity: O(n)
func LinearSearchSorted(data []int, value int) bool {
	for _, item := range data {
		if item == value {
			return true
		} else if item > value {
			return false
		}
	}
	return false
}
