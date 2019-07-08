package main

import (
	"algorithm/BubbleSort"
	"fmt"
)

func main() {
	var a = []int{6, 5, 3, 1, 8, 7, 2, 4}
	//fmt.Println(BucketSort.BucketSort(a))
	//a = []int{6, 5, 3, 1, 8, 7, 2, 4}
	//fmt.Println(SelectionSort.SelectionSort(a))
	//a = []int{6, 5, 3, 1, 8, 7, 2, 4}
	//fmt.Println(InsertionSort.InsertionSort(a))
	a = []int{6, 5, 3, 1, 8, 7, 2, 4}
	fmt.Println(BubbleSort.BubbleSort(a))
}
