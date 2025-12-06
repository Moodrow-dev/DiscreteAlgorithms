package Sort

/*
Временная сложность – O(n^2)
Для почти сортированных - O(n)
*/

func InsertionSort(slice []int) int {
	comparisonCount := 0
	for i := 1; i < len(slice); i++ {
		key := slice[i]
		j := i - 1

		for ; j >= 0; j-- {
			comparisonCount++
			if slice[j] < key {
				break
			}
			slice[j+1] = slice[j]
		}
		slice[j+1] = key
	}
	return comparisonCount
}
