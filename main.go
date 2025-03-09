package main

import (
	"fmt"
	"test/sorting"
)

func main() {
	arr := []int{5, 3, 8, 4, 2}
	fmt.Println("До сортировки:", arr)

	// InsertionSort меняет массив внутри

	sorted := sorting.QuickSort(arr)
	fmt.Println(sorted)

	// MergeSort возвращает новый массив
	arr = []int{5, 3, 8, 4, 2} // Сбрасываем массив
	arr = sorting.MergeSort(arr)

	fmt.Println("После MergeSort:", arr)
}
