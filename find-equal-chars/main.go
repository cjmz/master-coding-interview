package main

import "fmt"

func main() {
	arr1 := []string{"a", "b", "c", "x"}
	arr2 := []string{"u", "y", "p"}

	fmt.Printf("Return: %t\n", searchWithNestedArrays(arr1, arr2))
	fmt.Printf("Return: %+v\n", searchWithoutNestedArrays(arr1, arr2))

}

// O(a + b)
// Space Complexity = O(a)
func searchWithoutNestedArrays(arr1 []string, arr2 []string) bool {
	a1 := make(map[string]bool, len(arr1))

	for i := 0; i < len(arr1); i++ {
		if !a1[arr1[i]] {
			a1[arr1[i]] = true
		}
	}

	for j := 0; j < len(arr2); j++ {
		if a1[arr2[j]] {
			return true
		}
	}

	return false
}

// O(n^2)
// Space Complexity = O(1)
func searchWithNestedArrays(arr1 []string, arr2 []string) bool {
	r := false
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if arr1[i] == arr2[j] {
				r = true
				break
			}
		}
	}
	return r
}
