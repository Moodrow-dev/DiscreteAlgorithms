package Search

// LinearSearch - Линейный поиск, средняя сложность алгоритма O(n)
func LinearSearch(slice []int, target int) (int, int) {
	var comparisons int
	for _, j := range slice {
		comparisons++
		if j == target {
			//log.Println("Линейный поиск: Объект ", target, "находится по индексу ", j)
			return comparisons, j
		}
	}
	//log.Println("Линейный поиск: Объект не найден")
	return comparisons, -1
}
