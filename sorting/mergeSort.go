package sorting

import "fmt"

/*
📝 Merge Sort (Сортировка слиянием)
Merge Sort (сортировка слиянием) — это рекурсивный алгоритм, который использует принцип разделяй и властвуй.

🔍 Как работает алгоритм?
1️⃣ Разделение:
Разделяем массив пополам до тех пор, пока не останутся отдельные элементы (размером 1).
2️⃣ Слияние:
Начинаем обратно объединять подмассивы, сортируя их по ходу слияния.
3️⃣ Результат:
В конце объединяем все части в отсортированный массив.
📌 Пример работы Merge Sort
arr := []int{5, 3, 8, 4, 2}
Шаг 1: Разделение
[5, 3, 8, 4, 2]
     ↓ Разделяем
[5, 3]    [8, 4, 2]
 ↓  ↓      ↓  Разделяем
[5] [3]   [8] [4, 2]
               ↓
             [4] [2]
Шаг 2: Слияние (и сортировка)

[5] [3] → [3, 5]
[4] [2] → [2, 4]
[2, 4] [8] → [2, 4, 8]
Шаг 3: Финальное слияние

[3, 5]  [2, 4, 8] → [2, 3, 4, 5, 8]
✅ Результат: [2, 3, 4, 5, 8]

*/

func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Добавляем оставшиеся элементы
	result = append(result, left[i:]...)

	result = append(result, right[j:]...)
	fmt.Println(result)

	return result
}

// Рекурсивная функция сортировки
func MergeSort(arr []int) []int { // arr = []int{5, 3, 8, 4, 2}
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}
