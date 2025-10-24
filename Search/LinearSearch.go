package Search

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
