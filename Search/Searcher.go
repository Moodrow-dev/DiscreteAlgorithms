package Search

import (
	"fmt"
)

func Compare(slice []int, target int, search1 func([]int, int) int, search2 func([]int, int) int) {
	iter1 := search1(slice, target)
	iter2 := search2(slice, target)
	fmt.Println(slice)
	fmt.Println(iter1, "сравнений vs ", iter2, " сравнений")
	if iter1 < iter2 {
		fmt.Println("Первый алгоритм поиска эффективнее")
	} else if iter1 > iter2 {
		fmt.Println("Второй алгоритм поиска эффективнее")
	} else {
		fmt.Println("Алгоритмы одинаково эффективны")
	}
}
