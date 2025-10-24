package Sort

import "fmt"

func Compare(slice []int, sort1 func([]int) int, sort2 func([]int) int) {
	iter1 := sort1(slice)
	iter2 := sort2(slice)
	fmt.Println(slice)
	fmt.Println(iter1, "сравнений vs ", iter2, " сравнений")
	if iter1 < iter2 {
		fmt.Println("Первая сортировка эффективнее")
	} else if iter1 > iter2 {
		fmt.Println("Вторая сортировка эффективнее")
	} else {
		fmt.Println("Сортировки одинаково эффективны")
	}
}
