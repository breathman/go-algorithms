package search

func BinarySearchIterative(data []int, value int) bool {

	min := 0
	mid := 0
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
