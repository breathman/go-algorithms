package sort

func BubbleSort(s []int) []int {

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			var a int
			if s[i] < s[j] {
				a = s[i]
				s[i] = s[j]
				s[j] = a
			}
		}
	}

	return s
}
