package main

import "fmt"

/**
两数之和，lc 1
*/
func main() {
	arr1 := []int{2, 7, 8, 9}
	ints := add1(arr1, 9)
	fmt.Println("ints =", ints)
	fmt.Println("---------------------------")
	sum := twoSum(arr1, 9)
	fmt.Println("sum =", sum)
}

/**
双重遍历，使用 target - v，直接遍历数组，找到对应值
*/
func add1(arr []int, target int) []int {
	for i, v := range arr {
		for k := i + 1; k < len(arr); k++ {
			if target-v == arr[k] {
				return []int{i, k}
			}
		}
	}
	return []int{}
}

/**
使用map的key找到对应的下标
value, exist := m[target-k]; exist ==> 这个意思就是先复制，如果map中存在 target - k，对应的下标，则返回；
*/
func twoSum(nums []int, target int) []int {
	result := []int{}
	m := make(map[int]int)
	for i, k := range nums {
		if value, exist := m[target-k]; exist {
			result = append(result, value)
			result = append(result, i)
		}
		m[k] = i
	}
	return result
}
