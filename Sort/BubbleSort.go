package Sort

import "log"

type BubbleSort struct {
	Slice []int
	Sortable
}

// BubbleSort – пузырьковая сортировка, средняя сложность алгоритма O(n^2), в лучшем случае O(n)
func (bs BubbleSort) Sort() int {
	comparisonCount := 0
	slice := bs.Slice
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
	log.Println("Пузырьковая сортировка: ", slice)
	return comparisonCount
}

func (bs BubbleSort) GetSortComplexity() string {
	return "O(n^2), O(n) лучшая"
}

func (bs BubbleSort) GetSlice() []int {
	return bs.Slice
}
