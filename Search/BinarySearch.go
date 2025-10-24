package Search

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
