package Sort

func InsertionSort(slice []int) int {
	comparisonCount := 0
	for i := 1; i < len(slice); i++ {
		key := slice[i]
		j := i - 1

		for ; j >= 0 && key < slice[j]; j-- {
			slice[j+1] = slice[j]
			comparisonCount++
		}

		slice[j+1] = key
	}
	return comparisonCount
}
