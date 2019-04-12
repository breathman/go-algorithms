package sort

import (
	"github.com/breathman/algoritms/utils"
	"github.com/influxdata/influxdb/pkg/testing/assert"
	"testing"
)

const (
	dataRange  = 1000
	elemsCount = 1000
	runsCount  = 1000
)

func TestBubbleSort(t *testing.T) {
	s := []int{2, 5, 3, 8, 1, 1, 6}
	res := BubbleSort(s)
	assert.Equal(t, res, []int{1, 1, 2, 3, 5, 6, 8})
}

func BenchmarkBubbleSort(b *testing.B) {
	data := utils.RandSlice(dataRange, elemsCount)
	b.N = runsCount
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BubbleSort(data)
	}
}
