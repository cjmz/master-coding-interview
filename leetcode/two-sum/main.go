package main

import "fmt"

func main() {
	a := []int{2, 7, 11, 15}
	fmt.Println(twoSum2(a, 9))
}

// https://leetcode.com/problems/two-sum/

// Time Complexity: O(n)
// Space Complexity: O(1)
func twoSum2(nums []int, target int) []int {
	idx := make(map[int]int, len(nums))

	for i := 0; i <= len(nums)-1; i++ {
		s := target - nums[i]
		if ridx, ip := idx[s]; ip {
			return []int{ridx, i}
		}

		idx[nums[i]] = i
	}

	return []int{}
}

// Time Complexity: O(n^2)
// Space Complexity: O(1
func twoSum(nums []int, target int) []int {
	r := make([]int, 2)

	for i := 0; i <= len(nums)-1; i++ {
		for j := i + 1; j <= len(nums)-1; j++ {
			if nums[i]+nums[j] == target {
				r[0] = i
				r[1] = j
			}
		}
	}

	return r
}
