package main

import (
	"fmt"
	"github.com/breathman/algoritms/search"
)

var array = []int{4, 3, 3, 3, 2, 1, 4, 5, 2, 4, 9}
var item = 2

func main() {

	ok := search.LinearSearchUnsorted(array, 1)
	fmt.Println(ok)

	//res := quickSort(array)
	//res := binarySearch(item, array)

	//g.Init()
	//g.AddEdge("BTC", "ETH")
	//g.AddEdge("BTC", "DCH")
}
