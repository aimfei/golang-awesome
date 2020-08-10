package main

import "sort"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	len := len(nums)
	var newLen int = len
	for i, v := range nums {
		if i < len-1 {
			if v == nums[i+1] {
				newLen--
			}
		} else {
			if nums[i] == nums[i-1] {
				newLen--
			}
		}
	}
	return newLen
}
