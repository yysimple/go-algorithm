package main

import (
	"fmt"
	"sort"
)

/**
在数组中求出三个数符合 a + b + c = 0？ 返回对应的 abc的值，lc 15
总结：这里是采用双指针法：
	- 将数组排序
	- 先从0号索引开始固定值，然后 左指针 left 从下一个索引位置开始 i + 1, 右指针则从最后 len(nums) - 1开始
	- 从0开始，或者相邻两个索引对应的值不同（相同没有意义，相当于已经查询了一遍）才进行比对查值
	- 只要指针没有碰撞，则双指针继续“靠近”
		- 当三个值相加 = 0 则是符合的三个数，但是防止内层再次遍历，nums[left] = nums[left + 1]的情况下都会符合条件，所以直接跳过这次循环，直接给left+1
		- 三个值加起来 < 0 则 left + 1
		- 三个值加起来 > 0 则 right - 1
*/
func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	zero := threeNumAddToZero(nums)
	fmt.Println("zero =", zero)
}

func threeNumAddToZero(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		left, right := i+1, len(nums)-1
		if left >= len(nums) {
			break
		}

		if i == 0 || nums[i] != nums[i-1] {
			for left < right {
				target := 0 - nums[i]
				if nums[left]+nums[right] == target {
					num := []int{nums[i], nums[left], nums[right]}
					result = append(result, num)
					// 这里加了两层 for ，是为了避免取到的与 nums[i] 是同一个索引上面的值，然后就是往后顺位一个，这里可以打印索引看下具体情况，或者断点
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				} else if nums[left]+nums[right] < target {
					left++
				} else {
					right--
				}
			}
		}

	}
	return result

}
