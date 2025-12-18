package Sort

import "log"

type QuickSort struct {
	Slice []int
	Sortable
}

// QuickSort – быстрая сортировка, средняя сложность алгоритма O(n log n), в худшем случае O(n^2)
func (qs QuickSort) Sort() int {
	params := qs.quickSortParams(0, len(qs.Slice)-1)
	log.Println("Быстрая сортировка: ", qs.Slice)
	return params
}

func (qs QuickSort) quickSortParams(low, high int) int {
	comparisonCount := 0
	slice := qs.Slice

	if len(slice) == 0 || low >= high {
		return comparisonCount
	}

	border := slice[low]
	i, j := low, high

	for i <= j {
		for i <= high && slice[i] < border {
			i++
			comparisonCount++
		}
		for j >= low && slice[j] > border {
			j--
			comparisonCount++
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
		comparisonCount += qs.quickSortParams(low, j)
	}
	if high > i {
		comparisonCount += qs.quickSortParams(i, high)
	}
	return comparisonCount
}

func (qs QuickSort) GetSortComplexity() string {
	return "O(n log n), O(n^2) худшая"
}

func (qs QuickSort) GetSlice() []int {
	return qs.Slice
}
