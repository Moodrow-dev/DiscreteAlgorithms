package Search

import "log"

// LinearSearch - Линейный поиск, средняя сложность алгоритма O(n)
func LinearSearch(slice []int, target int) int {
	var comparisons int
	for _, j := range slice {
		comparisons++
		if j == target {
			log.Println("Линейный поиск: Объект ", target, "находится по индексу ", j)
			return comparisons
		}
	}
	log.Println("Линейный поиск: Объект не найден")
	return comparisons
}
