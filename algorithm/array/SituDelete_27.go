package main

import "fmt"

/**
原地删除，lc：27

*/
func main() {
	nums := []int{3, 2, 2, 3}
	arr := appendArr(nums, 3)
	fmt.Println("arr =", arr)
}

func appendArr(nums []int, val int) []int {
	for i := 0; i < len(nums); {
		if nums[i] == val {
			// 删除指定位置上的元素，这里删除了，i还是等于0，证明第一个元素符合删除的条件
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	return nums
}
