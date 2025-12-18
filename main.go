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

	//Sort.Compare(
	//	SliceTools.GenerateRandomSlice(10),
	//	Sort.InsertionBinarySort,
	//	Sort.InsertionSort,
	//)
	//Sort.Compare(
	//	SliceTools.GenerateSortedSlice(10),
	//	Sort.InsertionBinarySort,
	//	Sort.InsertionSort,
	//)
	//Sort.Compare(
	//	SliceTools.GenerateReversedSlice(10),
	//	Sort.InsertionBinarySort,
	//	Sort.InsertionSort,
	//)
	const SliceLength = 10

	slice := SliceTools.GenerateRandomSlice(SliceLength)
	//slice := SliceTools.GenerateSortedSlice(SliceLength)
	//slice := SliceTools.GenerateReversedSlice(SliceLength)

	//sort1 := Sort.InsertionBinarySort{Slice: append([]int{}, slice...)}
	//
	//sort2 := Sort.InsertionSort{Slice: append([]int{}, slice...)}

	sort1 := Sort.QuickSort{Slice: append([]int{}, slice...)}

	sort2 := Sort.MergeSort{Slice: append([]int{}, slice...)}

	Sort.Compare(sort1, sort2)
}
