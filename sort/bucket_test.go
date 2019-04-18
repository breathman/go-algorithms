package sort

import (
	"github.com/breathman/algoritms/utils"
	"gotest.tools/assert"
	"testing"
)

func TestBucketSort(t *testing.T) {
	s1 := []int{2, 5, 3, 8, 6, 4, 5}
	s2 := []int{2, 5, 3, 8, 1, 1, 6}
	res1 := BucketSort(s1, 2, 8)
	res2 := BucketSort(s2, 0, 10)
	assert.DeepEqual(t, res1, []int{2, 3, 4, 5, 5, 6, 8})
	assert.DeepEqual(t, res2, []int{1, 1, 2, 3, 5, 6, 8})

}

func BenchmarkBucketSort(b *testing.B) {
	data := utils.RandSlice(dataRange, elemsCount)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BucketSort(data, 0, dataRange)
	}
}
