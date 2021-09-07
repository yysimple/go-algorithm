package main

import "fmt"

/**
给定一个包含 0, 1, 2, ..., n 中 n 个数的序列，找出 0 .. n 中没有出现在序列中的那个数, lc 268
*/
func main() {
	nums := []int{0, 1, 2, 3, 4, 6}
	math := lostNumForMath(nums)
	fmt.Println("math: ", math)
	bit := lostNumForBit(nums)
	fmt.Println("lost num: ", bit)
}

/**
这是通过数学公式进行求和((n+1)n)/2
然后减去数组里面的值，进行相加
*/
func lostNumForMath(nums []int) int {
	numsLen := len(nums)
	total := numsLen * (numsLen + 1) / 2
	for _, v := range nums {
		total -= v
	}
	return total
}

/**
这里还是使用异或的思路来解题的，相消原理，这里因为是找从0开始 0...n，中间缺失的一个数
所以这里了利用索引的也行，不管顺序，0 ^ 1 ^ 2 ^ 0 ^ 1这种，最后肯定是会相消的；
然后因为最后一位无法相消，所以刚好数组长度刚好可以做到相消（因为不缺失，连续的话最后一位 应该是等于 len(nums) - 1的，因为缺失了一位，所以是len(nums)）
*/
func lostNumForBit(nums []int) int {
	res := 0
	for i, v := range nums {
		res ^= v ^ i
	}
	return res ^ len(nums)
}
