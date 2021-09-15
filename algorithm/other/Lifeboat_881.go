package main

import (
	"fmt"
	"sort"
)

/**
救生艇,lc 881
*/
func main() {
	nums := []int{3, 2, 2, 1, 5}
	i := findMinBoat(nums, 3)
	fmt.Println("num: ", i)
}

func findMinBoat(nums []int, limit int) int {
	num := 0
	left, right := 0, len(nums)-1
	sort.Ints(nums)
	for left <= right {
		if nums[left]+nums[right] <= limit {
			// 这里是当两者加起来 小于 limit，则两个人一起上船 两指针都向中间移动
			left++
			right--
		} else {
			// 如果超重，则“胖的”自己做一个船
			right--
		}
		num++
	}
	return num
}
