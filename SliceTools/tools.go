package SliceTools

import (
	"log"
	"math/rand"
)

var (
	ZeroItemSlice = []int{}  // Нулевой слайс
	OneItemSlice  = []int{1} // Единичный слайс
)

// GenerateRandomSlice генерирует случайный слайс с заданной длиной
func GenerateRandomSlice(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(size)
	}
	log.Println("Случайный массив:", slice)
	return slice
}

// GenerateSortedSlice генерирует заведомо отсортированный по возрастанию слайс с заданной длиной
func GenerateSortedSlice(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = i
	}
	log.Println("Сортированный массив: ", slice)
	return slice
}

// GenerateReversedSlice генерирует слайс отсортированный по убыванию
func GenerateReversedSlice(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = size - i - 1
	}
	log.Println("Сортированный по убыванию массив: ", slice)
	return slice
}
