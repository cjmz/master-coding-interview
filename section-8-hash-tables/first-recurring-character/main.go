package main

import "fmt"

func main() {
	a := []int{8, 12, 5, 1, 9, 12, 3, 4}
	fmt.Println(solve(a))
	fmt.Println(solve2(a))
}

// O(n^2)
// NO-HASH TABLES
func solve(n []int) int {
	// a := 2

	for i := 0; i < len(n); i++ {
		for j := i + 1; j < len(n); j++ {
			if n[i] == n[j] {
				return n[i]
			}
		}
	}

	return 0
}

// O(n)
// WITH HASH TABLES
func solve2(n []int) int {
	a := make(map[int]bool, len(n))

	for i := 0; i < len(n); i++ {
		if !a[n[i]] {
			a[n[i]] = true
			continue
		}

		return n[i]
	}

	return 0
}
