package Sort

import (
	"DiscreteAlgorithms/Search"
	"log"
)

// InsertionSort – сортировка вставками, средняя сложность алгоритма O(n^2), в лучшем случае O(n)
func InsertionSort(slice []int) int {
	comparisonCount := 0
	for i := 1; i < len(slice); i++ {
		key := slice[i]
		j := i - 1

		for ; j >= 0; j-- {
			comparisonCount++
			slice[j+1] = slice[j]
		}
		slice[j+1] = key
	}
	log.Println("Сортировка вставками: ", slice)
	return comparisonCount
}

// InsertionBinarySort – сортировка вставками с бинарным поиском, средняя сложность алгоритма O(log n)
func InsertionBinarySort(slice []int) int {
	comparisonCount := 0
	for i := 1; i < len(slice); i++ {
		key := slice[i]
		j := i - 1

		index := Search.BinarySearch(slice[0:j], key)

		for ; j >= index; j-- {
			comparisonCount++
			slice[j+1] = slice[j]
		}
		slice[j+1] = key
	}
	log.Println("Сортировка вставками с бинарным поиском: ", slice)
	return comparisonCount
}
