package main

import "fmt"

/**
旋转排序数组中的最小值Ⅰ,lc 153

*/
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	min := findMin(nums)
	fmt.Println("min = ", min)
}

/**
这里就是典型的二分法
*/
func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}
