package Sort

import "log"

/*
Временная сложность – O(n log n)
В худшем случае - O(n^2)
*/

// ShellSort – сортировка Шелла, средняя сложность алгоритма O(n log n), в худшем случае O(n^2)
func ShellSort(slice []int) int {
	comparisonCount := 0
	length := len(slice)
	step := length / 2

	for step > 0 {
		for i := 0; i < length-step; i++ {
			j := i
			for j >= 0 && slice[j] > slice[j+step] {
				slice[j], slice[j+step] = slice[j+step], slice[j]
				j--
				comparisonCount++
			}
			comparisonCount++

		}
		step /= 2
	}
	log.Println("Сортировка Шелла: ", slice)
	return comparisonCount
}
