package main

import "fmt"

/**
第k个最大元素,lc 215
*/
func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 5
	largest := findKthLargest(nums, k)
	fmt.Println("largest: ", largest)
}

func findKthLargest(nums []int, k int) int {
	idx := quickSort(0, len(nums)-1, len(nums)-k, nums)
	return nums[idx]
}

func quickSort(l, r, pos int, nums []int) int {
	povit_idx := partition(l, r, nums)
	if pos == povit_idx {
		return povit_idx
	} else if pos > povit_idx {
		return quickSort(povit_idx+1, r, pos, nums)
	} else {
		return quickSort(l, povit_idx-1, pos, nums)
	}
}

func partition(l, r int, nums []int) int {
	s := l
	povit_value := nums[l]
	for l < r {
		for l < r && nums[r] >= povit_value {
			r--
		}
		for l < r && nums[l] <= povit_value {
			l++
		}
		if l < r {
			nums[l] = nums[r] ^ nums[l]
			nums[r] = nums[l] ^ nums[r]
			nums[l] = nums[r] ^ nums[l]
		}
	}
	nums[s] = nums[l]
	nums[l] = povit_value
	return l
}
