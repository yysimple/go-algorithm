package main

import "fmt"

/**
最长上升子序列，lc 300
*/
func main() {
	nums := []int{1, 9, 5, 9, 3}
	lis := lengthOfLIS(nums)
	fmt.Println("arr =", lis)
}

func lengthOfLIS(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	newArr := make([]int, len(nums))
	res := 1
	for i := 0; i < len(nums); i++ {
		// 这个就是用来记录索引位置为多少的时候，最大连续上升区间是多少
		newArr[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				newArr[i] = max(newArr[i]+1, newArr[i])
			}
		}
		res = max(res, newArr[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
