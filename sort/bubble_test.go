package sort

import (
	"github.com/breathman/algoritms/utils"
	"gotest.tools/assert"
	"testing"
)

const (
	dataRange  = 10000
	elemsCount = 100
)

func TestBubbleSort(t *testing.T) {
	s := []int{2, 5, 3, 8, 1, 1, 6}
	res := BubbleSort(s)
	assert.DeepEqual(t, res, []int{1, 1, 2, 3, 5, 6, 8})
}

func BenchmarkBubbleSort(b *testing.B) {
	data := utils.RandSlice(dataRange, elemsCount)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BubbleSort(data)
	}
}
