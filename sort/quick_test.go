package sort

import (
	"github.com/breathman/algoritms/utils"
	"gotest.tools/assert"
	"testing"
)

func TestQuickSort(t *testing.T) {
	s := []int{2, 5, 3, 8, 1, 1, 6}
	res := QuickSort(s)
	assert.DeepEqual(t, res, []int{1, 1, 2, 3, 5, 6, 8})
}

func BenchmarkQuickSort(b *testing.B) {
	data := utils.RandSlice(dataRange, elemsCount)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		QuickSort(data)
	}
}
