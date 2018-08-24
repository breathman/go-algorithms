package search

import (
	"math/rand"
	"testing"
)

var udata = []int{92, 4, 2, 49, 47, 32, 90, 95, 35, 85, 41, 84, 82, 48, 6, 94, 69, 20, 50, 87, 79, 28, 67, 15, 36, 55, 75, 93, 91, 56, 8, 12, 42, 16, 66, 59, 57, 5, 33, 1, 54, 97, 80, 77, 65, 52, 29, 81, 14, 13, 86, 64, 98, 68, 46, 58, 76, 70, 11, 96, 60, 31, 34, 44, 30, 22, 61, 24, 43, 26, 62, 21, 74, 83, 39, 72, 73, 88, 45, 27, 0, 10, 89, 51, 18, 37, 17, 3, 40, 9, 19, 63, 71, 38, 53, 78, 7, 23, 99, 25}
var sdata = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99}

func TestLinearSearchUnsorted(t *testing.T) {
	var goal = 54
	res := LinearSearchUnsorted(udata, goal)

	if !res {
		t.Errorf("Elem %d not found", goal)
	}
}

func BenchmarkLinearSearchUnsorted(b *testing.B) {
	goal := rand.Intn(100)
	for i := 0; i < b.N; i++ {
		LinearSearchUnsorted(udata, goal)
	}
}

func TestLinearSearchSorted(t *testing.T) {
	var goal = 54
	res := LinearSearchSorted(sdata, goal)

	if !res {
		t.Errorf("Elem %d not found", goal)
	}
}

func BenchmarkLinearSearchSorted(b *testing.B) {
	goal := rand.Intn(100)
	for i := 0; i < b.N; i++ {
		LinearSearchSorted(sdata, goal)
	}
}
