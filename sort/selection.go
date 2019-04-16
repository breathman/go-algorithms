package sort

func SelectionSort(s []int) []int {

	for i := len(s) - 1; i > 0; i-- {
		maxIndex := i
		maxSubIndex := 0
		for j := 0; j < i-1; j++ {
			if s[j] < s[j+1] {
				maxSubIndex = j + 1
			}
		}
		if s[maxSubIndex] > s[maxIndex] {
			s[maxIndex], s[maxSubIndex] = s[maxSubIndex], s[maxIndex]
		}
	}

	return s
}
