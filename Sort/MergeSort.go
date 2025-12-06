package Sort

/*
Временная сложность – O(n log n)
В худшем случае - O(n^2)
*/

func MergeSort(slice []int) int {
	return mergeSortParams(slice, 0, len(slice)-1)
}

func mergeSortParams(slice []int, left, right int) int {
	comparisonCount := 0
	if left < right {
		mid := (left + right) / 2

		comparisonCount += mergeSortParams(slice, left, mid)
		comparisonCount += mergeSortParams(slice, mid+1, right)

		comparisonCount += merge(slice, left, mid, right)
	}
	return comparisonCount
}

func merge(slice []int, left, mid, right int) int {
	comparisonCount := 0

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
