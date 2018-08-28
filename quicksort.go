package main

func quickSort(array []int) []int {
	var (
		less   []int
		more   []int
		middle []int
	)

	if len(array) < 2 {
		return array
	} else {
		pivot := array[0]
		for _, item := range array[1:] {
			switch {
			case pivot > item:
				less = append(less, item)
			case pivot < item:
				more = append(more, item)
			case pivot == item:
				middle = append(middle, item)
			}
		}

		less, more = quickSort(less), quickSort(more)

		less = append(less, pivot)
		less = append(less, middle...)
		less = append(less, more...)

		return less
	}
}
