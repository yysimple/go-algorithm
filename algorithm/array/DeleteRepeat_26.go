package main

import "fmt"

/**
删除**排序数组**中的重复项，lc：26
返回移除后数组的新长度。不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成
*/
func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	repeat := deleteRepeat(nums)
	fmt.Println("repeat =", repeat)

}

/**
1. 定义当前元素 i 和下一个元素 i + 1
2. 索引i值 = i+1值时，删除i+1号元素
3. 不等时，i++，往后移动（因为时有序的）
4. 重复操作，当i+1指向最后一个元素时，跳出循环
*/
func deleteRepeat(nums []int) []int {
	for i := 0; i+1 < len(nums); {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	return nums
}
