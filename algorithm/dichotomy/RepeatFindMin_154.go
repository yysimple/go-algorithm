package main

import "fmt"

/**
有重复的数组，求 最小值 lc，154
*/
func main() {
	nums := []int{1, 2, 2, 4, 5, 6, 7, 8}
	min := repeatFindMin(nums)
	fmt.Println("min = ", min)
}

func repeatFindMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] < nums[right] {
			right = mid
		} else if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right--
		}
	}
	return nums[left]
}
