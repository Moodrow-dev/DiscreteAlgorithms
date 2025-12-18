package Sort

import (
	"log"
)

type BinarySort struct {
	Slice []int
	Sortable
}

// BinarySort - бинарная сортировка, средняя сложность алгоритма O(log n)
func (bs BinarySort) Sort() int {
	comparisonCount := 0

	slice := bs.Slice

	for i := 1; i < len(slice); i++ {
		key := slice[i]
		left, right := 0, i-1

		for left <= right {
			mid := left + (right-left)/2
			if key < slice[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
			comparisonCount++
		}
		for j := i; j > left; j-- {
			slice[j] = slice[j-1]
		}
		slice[left] = key
	}
	log.Println("Бинарная сортировка: ", slice)
	return comparisonCount
}

func (bs BinarySort) GetSortComplexity() string {
	return "O(log n)"
}

func (bs BinarySort) GetSlice() []int {
	return bs.Slice
}
