package search

import (
	"math/rand"
	"testing"
)

func TestBinarySearchIterative(t *testing.T) {
	var goal = 54
	res := BinarySearchIterative(sdata, goal)

	if !res {
		t.Errorf("Elem %d not found", goal)
	}
}

func BenchmarkBinarySearchIterative(b *testing.B) {
	goal := rand.Intn(100)
	for i := 0; i < b.N; i++ {
		BinarySearchIterative(sdata, goal)
	}
}

func TestBinarySearchRecursive(t *testing.T) {
	var goal = 54
	res := BinarySearchRecursive(sdata, goal, 0, len(sdata)-1)

	if !res {
		t.Errorf("Elem %d not found", goal)
	}
}

func BenchmarkBinarySearchRecursive(b *testing.B) {
	goal := rand.Intn(100)
	for i := 0; i < b.N; i++ {
		BinarySearchRecursive(sdata, goal, 0, len(sdata)-1)
	}
}
