package Sort

/*
Временная сложность – O(n log n)
В худшем случае - O(n^2)
*/

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
	return comparisonCount
}
