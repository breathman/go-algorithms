package search

import (
	"github.com/breathman/algoritms/utils"
	"sort"
	"testing"
)

func TestLinearSearchUnsorted(t *testing.T) {
	data := utils.RandSlice(dataRange, elemsCount)
	index := utils.RandInRange(0, elemsCount)
	res := LinearSearchUnsorted(data, data[index])

	if !res {
		t.Errorf("Elem %d not found", data[index])
	}
}

func BenchmarkLinearSearchUnsorted(b *testing.B) {
	data := utils.RandSlice(dataRange, elemsCount)
	index := utils.RandInRange(0, elemsCount)
	b.N = runsCount
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LinearSearchUnsorted(data, data[index])
	}
}

func TestLinearSearchSorted(t *testing.T) {
	data := utils.RandSlice(dataRange, elemsCount)
	index := utils.RandInRange(0, elemsCount)
	sort.Ints(data)
	res := LinearSearchSorted(data, data[index])

	if !res {
		t.Errorf("Elem %d not found", data[index])
	}
}

func BenchmarkLinearSearchSorted(b *testing.B) {
	data := utils.RandSlice(dataRange, elemsCount)
	index := utils.RandInRange(0, elemsCount)
	sort.Ints(data)
	b.N = runsCount
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LinearSearchSorted(data, data[index])
	}
}
