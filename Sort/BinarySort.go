package Sort

func BinarySort(slice []int) int {
	comparisonCount := 0

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
	return comparisonCount
}
