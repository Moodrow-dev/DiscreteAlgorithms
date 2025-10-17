package ShellAndBubble

import "fmt"

func shellSort(slice []int) int {
	comparisons := 0
	length := len(slice)
	step := length / 2

	for step > 0 {
		for i := 0; i < length-step; i++ {
			j := i
			for j >= 0 && slice[j] > slice[j+step] {
				slice[j], slice[j+step] = slice[j+step], slice[j]
				j--
				comparisons++
			}
			comparisons++

		}
		step /= 2
	}
	return comparisons
}

func bubbleSort(slice []int) int {
	comparisons := 0
	length := len(slice)

	for i := 0; i < length-1; i++ {
		swapped := false // Флаг, были ли перестановки на этом проходе

		for j := 0; j < length-i-1; j++ {
			comparisons++
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
				swapped = true // Была перестановка
			}
		}

		// Если не было ни одной перестановки - массив отсортирован
		if !swapped {
			break
		}
	}
	return comparisons
}

func Test() {
	//dataW := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	dataE := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	//data2 := []int{12, 11, 0, 13, 5, 6, 6, 6, 6}

	fmt.Println("Количество сравнений (Bubble): ", bubbleSort(dataE))
	fmt.Println("Отсортированный массив: ", dataE)

	fmt.Println("Количество сравнений (Shell): ", shellSort(dataE))
	fmt.Println("Отсортированный массив: ", dataE)
}
