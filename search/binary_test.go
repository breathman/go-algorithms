package search

import (
	"github.com/breathman/algoritms/utils"
	"sort"
	"testing"
)

const (
	dataRange  = 2000
	elemsCount = 1000
)

func TestBinarySearchIterative(t *testing.T) {
	data := utils.RandSlice(dataRange, elemsCount)
	index := utils.RandInRange(0, elemsCount)
	sort.Ints(data)
	res := BinarySearchIterative(data, data[index])

	if !res {
		t.Errorf("Elem %d not found", data[index])
	}
}

func BenchmarkBinarySearchIterative(b *testing.B) {
	data := utils.RandSlice(dataRange, elemsCount)
	index := utils.RandInRange(0, elemsCount)
	sort.Ints(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BinarySearchIterative(data, data[index])
	}
}

func TestBinarySearchRecursive(t *testing.T) {
	data := utils.RandSlice(dataRange, elemsCount)
	index := utils.RandInRange(0, elemsCount)
	sort.Ints(data)
	res := BinarySearchRecursive(data, data[index], 0, len(data)-1)

	if !res {
		t.Errorf("Elem %d not found", data[index])
	}
}

func BenchmarkBinarySearchRecursive(b *testing.B) {
	data := utils.RandSlice(dataRange, elemsCount)
	index := utils.RandInRange(0, elemsCount)
	sort.Ints(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BinarySearchRecursive(data, data[index], 0, len(data)-1)
	}
}
