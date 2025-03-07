package main

import (
	"fmt"
	"test/sorting"
)

func main() {
	arr := []int{5, 3, 8, 4, 2}
	fmt.Println("До сортировки:", arr)
	sorting.InsertionSort(arr)
	fmt.Println("После сортировки:", arr)
}
