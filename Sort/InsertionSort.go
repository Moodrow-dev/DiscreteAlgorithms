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

		// Сдвигаем элементы вправо, пока не найдем позицию для key
		for j >= 0 && slice[j] > key {
			comparisonCount++
			slice[j+1] = slice[j]
			j--
		}
		// Если цикл завершился по j < 0, мы не учли последнее сравнение
		if j >= 0 {
			comparisonCount++ // сравнение slice[j] > key, которое оказалось false
		}

		slice[j+1] = key
	}
	log.Println("Сортировка вставками: ", slice)
	return comparisonCount
}

// InsertionBinarySort – сортировка вставками с бинарным поиском, средняя сложность алгоритма O(n log n)
func InsertionBinarySort(slice []int) int {
	comparisonCount := 0
	for i := 1; i < len(slice); i++ {
		key := slice[i]

		// Бинарный поиск позиции для вставки
		//fmt.Println(slice[0:i], key)
		comparison, index := Search.BinarySearch(slice[0:i], key)
		comparisonCount += comparison

		if index < 0 {
			index = 0
		}

		// Если бинарный поиск возвращает позицию существующего элемента,
		// вставляем после него (для стабильности сортировки)
		// Ищем позицию после всех элементов <= key
		for index < i && slice[index] <= key {
			index++
		}

		// Сдвигаем элементы
		for j := i; j > index; j-- {
			slice[j] = slice[j-1]
		}

		slice[index] = key
	}
	log.Println("Сортировка вставками с бинарным поиском: ", slice)
	return comparisonCount
}
