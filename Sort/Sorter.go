package Sort

import "fmt"

//func Compare(slice []int, sort1 func([]int) int, sort2 func([]int) int) {
//	slice1 := make([]int, len(slice))
//	slice2 := make([]int, len(slice))
//	copy(slice1, slice)
//	copy(slice2, slice)
//	iter1 := sort1(slice1)
//	iter2 := sort2(slice2)
//	fmt.Println(slice)
//	fmt.Println(iter1, "сравнений vs ", iter2, " сравнений")
//	if iter1 < iter2 {
//		fmt.Println("Первая сортировка эффективнее")
//	} else if iter1 > iter2 {
//		fmt.Println("Вторая сортировка эффективнее")
//	} else {
//		fmt.Println("Сортировки одинаково эффективны")
//	}
//}

func Compare(s1, s2 Sortable) {
	fmt.Println("Сложность первой сортировки: ", s1.GetSortComplexity())
	fmt.Println("Сложность второй сортировка: ", s2.GetSortComplexity())

	slice := append([]int{}, s1.GetSlice()...)

	iter1 := s1.Sort()
	iter2 := s2.Sort()
	fmt.Println(fmt.Sprintf("Первоначальный слайс: %v", slice))
	fmt.Println(iter1, "сравнений vs ", iter2, " сравнений")
	if iter1 < iter2 {
		fmt.Println("Первая сортировка эффективнее")
	} else if iter1 > iter2 {
		fmt.Println("Вторая сортировка эффективнее")
	} else {
		fmt.Println("Сортировки одинаково эффективны")
	}

}

type Sortable interface {
	GetSortComplexity() string
	Sort() int
	GetSlice() []int
}
