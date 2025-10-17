package QuickAndMerge

import "fmt"

func QuickSort(slice []int, low, high int) int {
	comparisons := 0

	if len(slice) == 0 || low >= high {
		return comparisons
	}

	border := slice[low]
	i, j := low, high

	for i <= j {
		for i <= high && slice[i] < border {
			i++
			comparisons++
		}
		for j >= low && slice[j] > border {
			j--
			comparisons++
		}

		if i <= j {
			if i != j {
				slice[i], slice[j] = slice[j], slice[i]
			}
			i++
			j--
		}
	}

	if low < j {
		comparisons += QuickSort(slice, low, j)
	}
	if high > i {
		comparisons += QuickSort(slice, i, high)
	}
	return comparisons
}

func QuickSortWrapper(slice []int) int {
	return QuickSort(slice, 0, len(slice)-1)
}

func mergeSort(slice []int, left, right int) int {
	comparisons := 0
	if left < right {
		mid := (left + right) / 2

		comparisons += mergeSort(slice, left, mid)
		comparisons += mergeSort(slice, mid+1, right)

		comparisons += merge(slice, left, mid, right)
	}
	return comparisons
}

func merge(slice []int, left, mid, right int) int {
	comparisons := 0

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
		comparisons++
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
	return comparisons
}

func MergeSort(slice []int) int {
	return mergeSort(slice, 0, len(slice)-1)
}

func Test() {
	dataSort := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	dataSort2 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	fmt.Println("Исходный массив:", dataSort)
	fmt.Println("Количество сравнений (Quick Sort):", QuickSortWrapper(dataSort))
	fmt.Println("Отсортированный массив (Quick):", dataSort)

	fmt.Println("Количество сравнений (Merge Sort):", MergeSort(dataSort2))
	fmt.Println("Отсортированный массив (Merge):", dataSort2)
}
