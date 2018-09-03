package main

import (
	"fmt"
	"github.com/breathman/algoritms/random"
)

func main() {

	a := random.RandInRange(1, 100)
	fmt.Println(a)

	b := random.RandSliceInRange(10, 20, 20)
	fmt.Println(b)

	c := random.RandSlice(10, 20)
	fmt.Println(c)

}
