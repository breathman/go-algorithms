package sort

func MergeSort(ar []int) []int {
	res := make([]int, len(ar), len(ar))
	mergeSort(ar, res, 0, len(ar)-1)
	return ar
}

func mergeSort(ar, res []int, lIndex, rIndex int) {
	if lIndex >= rIndex {
		return
	}
	mid := (rIndex + lIndex) / 2
	mergeSort(ar, res, lIndex, mid)
	mergeSort(ar, res, mid+1, rIndex)
	merge(ar, res, lIndex, mid, rIndex)
}

func merge(ar, res []int, lIndex, mid, rIndex int) {
	lStart := lIndex
	lStop := mid
	rStart := mid + 1
	rStop := rIndex
	i := lIndex

	for lStart <= lStop && rStart <= rStop {
		if ar[lStart] < ar[rStart] {
			res[i] = ar[lStart]
			lStart++
		} else {
			res[i] = ar[rStart]
			rStart++
		}
		i++
	}

	for lStart <= lStop {
		res[i] = ar[lStart]
		lStart++
		i++
	}
	for rStart <= rStop {
		res[i] = ar[rStart]
		rStart++
		i++
	}
	for i := lIndex; i <= rIndex; i++ {
		ar[i] = res[i]
	}
}
