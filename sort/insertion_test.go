package sort

import (
	"github.com/breathman/algoritms/utils"
	"github.com/influxdata/influxdb/pkg/testing/assert"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	s := []int{2, 5, 3, 8, 1, 1, 6}
	res := InsertionSort(s)
	assert.Equal(t, res, []int{1, 1, 2, 3, 5, 6, 8})
}

func BenchmarkInsertionSort(b *testing.B) {
	data := utils.RandSlice(dataRange, elemsCount)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InsertionSort(data)
	}
}
