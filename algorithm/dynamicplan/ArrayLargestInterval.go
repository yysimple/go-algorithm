package main

import "fmt"

/**
最大子序和，lc 53
*/
func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4, 7}
	array := simpleSubMax(nums)
	fmt.Println("maxSub =", array)
}

/**
这个比较好理解，拿一个中间变量去接收每次相加的值：
	1. 相加之后小于0，那么res = 当前数组遍历索引位置的值
	2. 记录这个值到一个新的数组，然后继续往下遍历，重复这两个逻辑
	3. 最后遍历新的数组即可，取出最大值
*/
func maxSubArray(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	res := 0
	newArray := make([]int, len(nums))
	res = nums[0]
	for i := 1; i < len(nums); i++ {
		if res < 0 {
			res = nums[i]
		} else {
			newArray = append(newArray, res)
			res += nums[i]
		}
	}
	for i := 0; i < len(newArray); i++ {
		res = maxNum(res, newArray[i])
	}
	return res
}

/**
这个是直接用数组方式，不需要中间遍历的方式
*/
func maxSubArr(nums []int) int {
	if len(nums) < 0 {
		return 0
	}
	newArr := make([]int, len(nums))
	newArr[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if newArr[i-1] < 0 {
			newArr[i] = nums[i]
		} else {
			newArr[i] = newArr[i-1] + nums[i]
		}
	}
	res := 0
	for _, k := range newArr {
		res = maxNum(res, k)
	}
	return res
}

/**
这种方式就是记录每次的最大值，最后将res返回就行
*/
func simpleSubMax(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	newArr := make([]int, len(nums))
	res := nums[0]
	newArr[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		newArr[i] = maxNum(newArr[i-1]+nums[i], nums[i])
		res = maxNum(res, newArr[i])
	}
	return res
}

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}
