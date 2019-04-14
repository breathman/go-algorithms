package sort

func BubbleSort(s []int) []int {

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if s[i] < s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}

	return s
}
