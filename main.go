package main

import (
	"fmt"

	sortingAlgos "github.com/guitarkeegan/cs-algos/sorting-algos"
)

func main() {
	slice := []int{4, 2, 1, 3}
	sortingAlgos.BubbleSort(slice)
	fmt.Printf("Bubble Sort: %v\n", slice)

	slice = []int{4, 2, 1, 3}
	sortingAlgos.BubbleSort(slice)
	fmt.Printf("Insertion Sort: %v\n", slice)

	slice = []int{4, 2, 1, 3}
	slice = sortingAlgos.MergeSort(slice)
	fmt.Printf("Merge Sort: %v\n", slice)
}
