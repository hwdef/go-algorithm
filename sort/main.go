package main

import (
	"fmt"
	"sort/algorithm"
)

func main() {
	var a = []int{6, 5, 3, 1, 8, 7, 2, 4}
	algorithm.InsertionSort(a)
	fmt.Println(a)
}
