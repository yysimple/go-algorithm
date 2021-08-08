package main

import "fmt"

/**
最大子序和，lc 53
*/
func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4, 7}
	array := maxSubArr(nums)
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
	var newArray []int
	var indexArr []int
	res = nums[0]
	for i := 1; i < len(nums); i++ {
		if res < 0 {
			res = nums[i]
		} else {
			res += nums[i]
			newArray = append(newArray, res)
			indexArr = append(indexArr, i+1)
		}
	}
	fmt.Println("maxArr =", newArray)
	fmt.Println("indexArr =", indexArr)
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
		/**
		这里再来总结下动态规划的思想：
		首先我们的需求：给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
		这里需求解读一下，我们需要找到最大和的连续子数组,返回最大值；
		我们应该想到的是，拆分数组，每次都只是简单的对比前后相邻的值，然后取最大值存入到新的数组中；
		这里就是理解为根据需求，拆分成一个个小的逻辑单元，他们都有一定的关系（取最大值也就是代码中的 if-else），然后在聚合这些类似操作
		得到最终想要的，这里其实也就是按部就班的操作，但是又没有明确的第一步应该做什么，第二部应该做什么？只要能拆分成相同的逻辑单元就行，
		而这个逻辑单元就是状态转移方程；将连续子数组最大和，拆分成两数之间的最大值，只是后面一直在累加而已之前的最大值，然后又跟下一位对比求最大值；
		补充：还有就是动态规划一般需要另外一个数组进行辅助，记录下每次对比（最小逻辑单元处理）的结果，用来做最后的结果返回；
		*/
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
	fmt.Println("maxArr =", newArr)
	return res
}

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}
