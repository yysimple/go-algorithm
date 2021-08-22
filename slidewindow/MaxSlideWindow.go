package main

import (
	"fmt"
	"math"
)

/**
滑动窗口最大值，lc 239
*/
func main() {
	arr := []int{1, 3, -1, -3, 5, 3, 6, 7}
	window := findMaxWindow(arr, 3)
	fmt.Println("max window =", window)
	slidingWindow := maxSlidingWindow(arr, 3)
	fmt.Println("slidingWindow =", slidingWindow)

}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	//用切片模拟一个双端队列
	var queue []int
	var result []int
	for i := range nums {
		for i > 0 && (len(queue) > 0) && nums[i] > queue[len(queue)-1] {
			//将比当前元素小的元素祭天
			queue = queue[:len(queue)-1]
		}
		//将当前元素放入queue中
		queue = append(queue, nums[i])
		if i >= k && nums[i-k] == queue[0] {
			//维护队列，保证其头元素为当前窗口最大值
			queue = queue[1:]
		}
		if i >= k-1 {
			//放入结果数组
			result = append(result, queue[0])
		}
	}
	return result
}

func findMaxWindow(nums []int, winLen int) []int {
	res := make([]int, len(nums)-winLen+1)
	if len(nums) == 0 {
		return res
	}
	for i := 0; i < len(nums)-winLen+1; i++ {
		max := math.MinInt64
		for j := i; j < i+winLen; j++ {
			max = maxLen(max, nums[j])
		}
		res[i] = max
	}
	return res
}

func maxLen(a, b int) int {
	if a > b {
		return a
	}
	return b
}
