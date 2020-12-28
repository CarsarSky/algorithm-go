package main

import (
	"algorithm-go/sort"
	"fmt"
)

func main() {

	array := []int{1, 8, 3, 4, 9}
	b := sort.QuickSort(array, 0, len(array)-1)
	fmt.Print(b)
}
