package main

import "fmt"

/**
和为s的连续正数序列，单个值不输出，比如 [123456789] 输出 [234] [45]，而[9]是不输出的
*/
func main() {
	sequence := findContinuousSequence(9)
	fmt.Println("sequence =", sequence)
}

func findContinuousSequence(target int) [][]int {
	result := make([][]int, 0)
	// 左窗口，这里之后对应的是数组坐标，但是因为 连续正整数是从 1 开始的，所以这里初始化值为 1
	left := 1
	// 右窗口
	right := 1
	// 窗口值
	winSum := 0
	// 初始化连续正整数
	arr := make([]int, target)
	for i := range arr {
		arr[i] = i + 1
	}
	// 然后开始找子串，这里当left 超过了一半，就不能找到子串的 想想 是不是如此，比如 9 中间就是 9/2=4，如果left = 5 后面就不会有比 5 小的了
	for left <= target/2 {
		// 如果窗口字串值，小于目标值，窗口 右移
		if winSum < target {
			// 这里其实是进行累加操作，然后窗口 right 移动
			winSum += right
			right++
			// 这里数小于目标值，证明窗口现在的值过大，那从左边开始 -- ，这里也是个常识，如果一串数字过大，我们可以从小的数字，慢慢减去
		} else if winSum > target {
			// 这里是减小窗口长度，将窗口里的较小值去掉,然后左窗口边界 往右移动
			winSum -= left
			left++
			// 相等的情况，证明找到了，截取
		} else {
			// 这里上面也说了，是为了模拟 1 开始，所以这的索引其实是往后移了一位的
			result = append(result, arr[left-1:right-1])
			// 后面这两步，是为了继续往后找其他子串，然后窗口的值相应减小，左窗口边界也右移一位
			winSum -= left
			left++
		}
	}
	return result
}
