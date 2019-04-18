package sort

func BucketSort(ar []int, minRange, maxRange int) []int {
	size := maxRange - minRange + 1
	bucketAr := make([]int, size, size)

	for i := 0; i < len(ar); i++ {
		bucketAr[ar[i]-minRange]++
	}

	k := 0
	for i := 0; i < len(bucketAr); i++ {
		if bucketAr[i] != 0 {
			for j := 1; j <= bucketAr[i]; j++ {
				ar[k] = i + minRange
				k++
			}
		}
	}

	return ar
}
