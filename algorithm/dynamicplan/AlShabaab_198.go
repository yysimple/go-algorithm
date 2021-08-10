package main

import "fmt"

/**
打家劫舍，lc 198
*/
func main() {
	nums := []int{2, 7, 9, 3, 1, 2, 5, 6}
	i := rob(nums)
	fmt.Println("arr =", i)
}

/**
这个也比较简单，也是累加思想，因为要相隔一步，所以在选择下一个要累加的值，则需要关心下一中情况的最大值，其实这里是反向确定的：
	- i = 2的时候，逆推思想，只需要关心，a = nums[1] + nums[3] 和 b = nums[2]选最大值即可，明显 a 更大，所以此时 nums[2] = 11
	- i = 3的时候，a = nums[2] + nums[4] = 11 + 1  /  b = nums[3] = 3 所以此时 nums[3] = 12 一次类推，最后一个元素就是要取的值；
*/
func rob(nums []int) int {
	if len(nums) < 1 {
		return nums[0]
	}
	if len(nums) < 2 {
		return max(nums[0], nums[1])
	}
	nums[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		nums[i] = max(nums[i-2]+nums[i], nums[i-1])
	}
	return nums[len(nums)-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
