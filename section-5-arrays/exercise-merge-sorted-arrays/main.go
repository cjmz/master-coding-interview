package main

import (
	"fmt"
)

func main() {
	arr1 := []int{0, 3, 31, 68}
	arr2 := []int{4, 30, 33}

	fmt.Println(mergeSortedArrays(arr1, arr2))
}

func mergeSortedArrays(arr1 []int, arr2 []int) []int {
	a := arr1

	for i := 0; i <= len(arr2)-1; i++ {
		a = append(a, arr2[i])
	}

	// for j := len(a) - 1; j <= 0; j-- {
	// 	if j > 0 && a[j-1] > a[j] {
	// 		n := a[j]
	// 		a[j] = a[j+1]
	// 		a[j-1] = n
	// 	}
	// }

	return a
}

func sort(a []int) {
	for i := 1; i <= len(a)-1; i++ {

		for j := i; j >= 0; j-- {
			if a[i] < a[j] {
				a[i] = a[j]
			}
		}
	}
}

// [0, 3, 31, 68, 4, 30, 33]
