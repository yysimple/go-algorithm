package main

import "fmt"

/**
只出现一次的数字， lc 136
*/
func main() {
	nums := []int{2, 1, 2, 3, 5, 1, 3}
	one_01 := findOnlyOne_01(nums)
	fmt.Println("res：", one_01)
}

/**
这里是用到了 ^ 运算，如果相同则为0，不同为1，所以如果两个数相等的 ^ 运算，最后结果为 1
*/
func findOnlyOne_01(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res = res ^ nums[i]
	}
	return res
}
