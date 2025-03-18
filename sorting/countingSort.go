package sorting

func CountingSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}

	count := make([]int, max+1)

	for _, num := range arr {
		count[num]++
	}

	index := 0
	for num, freq := range count {
		for freq > 0 {
			arr[index] = num
			index++
			freq--
		}
	}

	return arr
}
