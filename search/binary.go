package search

// Complexity: O(log n)
func BinarySearchIterative(data []int, value int) bool {
	var min, mid int
	max := len(data) - 1

	for min <= max {
		mid = min + (max-min)/2
		if value == data[mid] {
			return true
		} else if value < data[mid] {
			max = mid - 1
		} else {
			min = mid + 1
		}
	}
	return false
}

// Complexity: O(log n)
func BinarySearchRecursive(data []int, value, min, max int) bool {
	mid := min + (max-min)/2
	if data[mid] == value {
		return true
	} else if value < data[mid] {
		return BinarySearchRecursive(data, value, min, mid-1)
	} else {
		return BinarySearchRecursive(data, value, mid+1, max)
	}

	return false
}
