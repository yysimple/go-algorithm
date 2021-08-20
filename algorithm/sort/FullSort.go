package main

import "fmt"

/**
全排序
*/
func main() {
	arr := []int{1, 2, 3}
	permutation(arr, 0, len(arr))
}

func permutation(arr []int, from, to int) {
	if from == to {
		fmt.Println(arr)
		return
	}
	first := arr[from]
	for i := from; i < to; i++ {
		itemRaw := arr[i]

		// 将i替换成first
		arr[i] = first
		arr[from] = itemRaw
		// 后面的元素全排
		permutation(arr, from+1, to)
		// 还原
		arr[from] = first
		arr[i] = itemRaw
	}
}

func permutation2(arr []int, cur, n int) {
	if cur == n {
		fmt.Println(arr)
		return
	}
	for i := 1; i <= n; i++ {
		// 判断在之前是否存在，不存在则使用
		ok := true
		for j := 0; j < cur; j++ {
			if i == arr[j] {
				ok = false
				break
			}
		}
		if ok {
			arr[cur] = i
			// fmt.Println(i)
			// 找下一个位置
			permutation2(arr, cur+1, n)
		}
	}
}
