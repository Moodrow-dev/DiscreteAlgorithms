package ConsoleSearchAl

import "fmt"

func LinearSearch(slice []int, target int) int {
	var comparisons int
	for _, j := range slice {
		comparisons++
		if j == target {
			return comparisons
		}
	}
	return comparisons
}

func BinarySearch(slice []int, target int) int {
	var comparisons int
	left, right := 0, len(slice)-1

	for left <= right {
		comparisons++
		mid := left + (right-left)/2

		if slice[mid] == target {
			return mid
		} else if slice[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return comparisons
}

func Test() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 7

	linearComparisons := LinearSearch(slice, target)
	fmt.Println("Линейный поиск: ", linearComparisons, " сравнений")
	binaryComparisons := BinarySearch(slice, target)
	fmt.Println("Бинарный поиск: ", binaryComparisons, " сравнений")
}
