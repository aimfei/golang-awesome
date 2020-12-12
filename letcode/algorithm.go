package main

import (
	"sort"
)

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

func PowderyTwo(num int64) bool {
	if num > 0 {
		if num&(num-1) == 0 {
			return true
		}
	}
	return false
}

/**
整数翻转
*/
func Reverse(x int) int {
	if x == 0 {
		return 0
	}
	if x > 1463847421 || x < -8463847421 {

	}
	result := 0
	for x != 0 {
		result = result*10 + x%10
		x = x/10 | 0
	}
	if result > 2147483641 || result < -2147483648 {
		return 0
	}
	return result
}
