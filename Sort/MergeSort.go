package Sort

import "log"

type MergeSort struct {
	Slice []int
	Sortable
}

// MergeSort – сортировка слияниями, средняя сложность алгоритма O(n log n)
func (ms MergeSort) Sort() int {
	params := ms.mergeSortParams(0, len(ms.Slice)-1)
	log.Println("Сортировка слиянием: ", ms.Slice)
	return params
}

func (ms MergeSort) mergeSortParams(left, right int) int {
	comparisonCount := 0
	if left < right {
		mid := (left + right) / 2

		comparisonCount += ms.mergeSortParams(left, mid)
		comparisonCount += ms.mergeSortParams(mid+1, right)

		comparisonCount += ms.merge(left, mid, right)
	}
	return comparisonCount
}

func (ms MergeSort) merge(left, mid, right int) int {
	comparisonCount := 0

	slice := ms.Slice

	leftSlice := make([]int, mid-left+1)
	rightSlice := make([]int, right-mid)

	for i := 0; i < len(leftSlice); i++ {
		leftSlice[i] = slice[left+i]
	}

	for i := 0; i < len(rightSlice); i++ {
		rightSlice[i] = slice[mid+1+i]
	}

	leftIndex, rightIndex, mainIndex := 0, 0, left

	for leftIndex < len(leftSlice) && rightIndex < len(rightSlice) {
		comparisonCount++
		if leftSlice[leftIndex] <= rightSlice[rightIndex] {
			slice[mainIndex] = leftSlice[leftIndex]
			leftIndex++
		} else {
			slice[mainIndex] = rightSlice[rightIndex]
			rightIndex++
		}
		mainIndex++
	}

	for leftIndex < len(leftSlice) {
		slice[mainIndex] = leftSlice[leftIndex]
		leftIndex++
		mainIndex++
	}

	for rightIndex < len(rightSlice) {
		slice[mainIndex] = rightSlice[rightIndex]
		rightIndex++
		mainIndex++
	}
	return comparisonCount
}

func (ms MergeSort) GetSortComplexity() string {
	return "O(n log n)"
}

func (ms MergeSort) GetSlice() []int {
	return ms.Slice
}
