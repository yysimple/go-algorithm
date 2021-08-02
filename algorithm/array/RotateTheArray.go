package main

import "fmt"

/**
旋转数组，lc 189
总结：其实就是前后颠倒，然后移动位数的元素，在此颠倒顺序即可

*/
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	index := changeArrIndex(arr, 3)
	fmt.Println("index =", index)
	ints := rotate(arr, 2)
	fmt.Println("ints =", ints)

}

/**
这种方式，就是互换位置，将移动位数 changeNum 作为元素最后几位索引对应的元素  len(arr)-changeNum+i
*/
func changeArrIndex(arr []int, changeNum int) []int {
	if changeNum > len(arr) {
		return arr
	}
	// 往指定位置上插入元素
	for i := 0; i < changeNum; i++ {
		arr = append(arr[:i], append([]int{arr[len(arr)-changeNum+i]}, arr[i:]...)...)
	}
	// 截取最新的数组
	arr = append(arr[:(len(arr) - changeNum)])
	return arr
}

/**
这个也差不多，先是全部调换顺序，然后在指定需要变换的位数的那些元素互换位置，然后剩下的后移的元素换下位置
*/
func rotate(nums []int, k int) []int {
	// 7 6 5 4 3 2 1
	reverse(nums)
	// 6 7
	reverse(nums[:k%len(nums)])
	// 1 2 3 4 5
	reverse(nums[k%len(nums):])
	return nums
}

func reverse(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}
