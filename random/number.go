package random

import (
	"math/rand"
	"time"
)

func RandInRange(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

func RandSlice(r, l int) []int {
	d := make([]int, l)
	for i := 0; i < l; i++ {
		rand.Seed(time.Now().UnixNano())
		d[i] = rand.Intn(r)
	}
	return d
}

func RandSliceInRange(min, max, l int) []int {
	s := make([]int, l)
	for i := 0; i < l; i++ {
		s[i] = RandInRange(min, max)
	}
	return s
}
