package sorting

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[len(arr)/2] // Выбираем опорный элемент
	left := []int{}
	right := []int{}

	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		} else if v > pivot {
			right = append(right, v)
		}
	}

	// Рекурсивно сортируем левую и правую части и собираем результат
	return append(append(QuickSort(left), pivot), QuickSort(right)...)
}
