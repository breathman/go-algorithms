package search

import (
	"math/rand"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	var goal = 54
	res := BinarySearchIterative(sdata, goal)

	if !res {
		t.Errorf("Elem %d not found", goal)
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	goal := rand.Intn(100)
	for i := 0; i < b.N; i++ {
		BinarySearchIterative(sdata, goal)
	}
}
