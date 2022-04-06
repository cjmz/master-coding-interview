package main

import "fmt"

type char struct {
	char  string
	found bool
}

func main() {
	arr1 := []string{"a", "b", "c"}
	arr2 := []string{"x", "y", "u", "a"}

	// fmt.Printf("Return: %t\n", searchWithNestedArrays(arr1, arr2))
	fmt.Printf("Return: %+v\n", searchWithoutNestedArrays(arr1, arr2))

}

func searchWithoutNestedArrays(arr1 []string, arr2 []string) []char {
	var s1 []char

	for i := 0; i < len(arr1); i++ {
		s1 = append(s1, char{char: arr1[i], found: true})
	}

	for j := 0; j < len(arr2); j++ {

	}

	return s1
}

// O(n^2)
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
