package InsertionBinarySort

import "fmt"

func InsertionBinarySort(slice []int) {
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
	fmt.Println("Количество сравнений (бинарная сортировка): ", comparisonCount)
}

func InsertionSort(slice []int) {
	comparisonCount := 0
	for i := 1; i < len(slice); i++ {
		key := slice[i]
		j := i - 1

		for ; j >= 0 && key < slice[j]; j-- {
			slice[j+1] = slice[j]
			comparisonCount++
		}

		slice[j+1] = key
	}
	fmt.Println("Количество сравнений (сортировка вставками): ", comparisonCount)
}

func Test() {
	//data := []int{12,11,0,13,5,6,6,6,6}
	dataW := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	data1 := make([]int, len(dataW))
	copy(data1, dataW)
	InsertionSort(dataW)
	fmt.Println("Отсортированный массив (сортировка вставками): ", dataW)

	InsertionBinarySort(data1)
	fmt.Println("Отсортированный массив (бинарная сортировка): ", data1)

}
