package Sort

func BubbleSort(slice []int) int {
	comparisonCount := 0
	length := len(slice)

	for i := 0; i < length-1; i++ {
		swapped := false // Флаг, были ли перестановки на этом проходе

		for j := 0; j < length-i-1; j++ {
			comparisonCount++
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
				swapped = true // Была перестановка
			}
		}

		// Если не было ни одной перестановки - массив отсортирован
		if !swapped {
			break
		}
	}
	return comparisonCount
}
