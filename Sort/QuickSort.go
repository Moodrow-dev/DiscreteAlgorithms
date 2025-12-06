package Sort

/*
Временная сложность – O(n log n)
В худшем случае - O(n^2)
*/

func QuickSort(slice []int) int {
	return quickSortParams(slice, 0, len(slice)-1)
}

func quickSortParams(slice []int, low, high int) int {
	comparisonCount := 0

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
		comparisonCount += quickSortParams(slice, low, j)
	}
	if high > i {
		comparisonCount += quickSortParams(slice, i, high)
	}
	return comparisonCount
}
