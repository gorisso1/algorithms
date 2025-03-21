package main

import "fmt"

func main() {
	arr1 := []int{1, 3, 5, 7}
	arr2 := []int{2, 4, 6, 8}
	sort := task1(arr1, arr2)
	fmt.Println(sort)

}

func task1(arr1, arr2 []int) []int {
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
