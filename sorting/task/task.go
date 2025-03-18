package sorting

import "fmt"

func Task1Bubble(arr1, arr2 []int) []int {
	mergedArr := append(arr1, arr2...)
	fmt.Println(mergedArr)
	for i := 0; i < len(mergedArr)-1; i++ {
		swapped := false
		for j := 0; j < len(mergedArr)-i-1; j++ {
			if mergedArr[j] > mergedArr[j+1] {
				mergedArr[j], mergedArr[j+1] = mergedArr[j+1], mergedArr[j]
				swapped = true
			}
		}
		if swapped == false {
			return mergedArr
		}
	}
	return mergedArr
}
