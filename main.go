package main

import (
	"fmt"
	"github.com/breathman/algoritms/utils"
)

func main() {

	a := utils.RandInRange(1, 100)
	fmt.Println(a)

	b := utils.RandSliceInRange(10, 20, 20)
	fmt.Println(b)

	c := utils.RandSlice(10, 20)
	fmt.Println(c)

}
