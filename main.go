package main

import (
	"DiscreteAlgorithms/SliceTools"
	"DiscreteAlgorithms/Sort"
)

func main() {
	//searchSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//target := 7
	//Search.Compare(searchSlice, target, Search.BinarySearch, Search.LinearSearch)
	//
	//sortSlice := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	//
	//Sort.Compare(sortSlice, Sort.QuickSort, Sort.BinarySort)

	Sort.Compare(
		SliceTools.GenerateRandomSlice(10),
		Sort.InsertionBinarySort,
		Sort.InsertionSort,
	)
	Sort.Compare(
		SliceTools.GenerateSortedSlice(10),
		Sort.InsertionBinarySort,
		Sort.InsertionSort,
	)
	Sort.Compare(
		SliceTools.GenerateReversedSlice(10),
		Sort.InsertionBinarySort,
		Sort.InsertionSort,
	)
}
