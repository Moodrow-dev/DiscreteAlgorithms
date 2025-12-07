package Search

import "log"

// BinarySearch – двоичный поиск, средняя сложность алгоритма O(log n)
func BinarySearch(slice []int, target int) int {
	var comparisons int
	left, right := 0, len(slice)-1

	for left <= right {
		mid := left + (right-left)/2
		comparisons++

		if slice[mid] == target {
			log.Println("Бинарный поиск: Объект ", target, " находится по индексу ", mid)
			return comparisons // индекс и количество сравнений
		} else if slice[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	log.Println("Бинарный поиск: Объект не найдем")
	return comparisons // не найден и количество сравнений
}
