package main

import (
	"fmt"
	"sort"
)

func main() {
	// 无序的两个数组
	arr1 := []int{4, 9, 4}
	arr2 := []int{1, 9, 5, 4, 9}
	forMap := findForMap(arr1, arr2)
	fmt.Println("无序交集：===> ", forMap)
	fmt.Println("---------------------------")
	forSorted := findForSorted(arr1, arr2)
	fmt.Println("有序交集：===> ", forSorted)
}

/**
	首先这是无序的数组求交集，大概思路就是
	1. 新建一个map
	2. 以其中一个数组 arr1 遍历，将其值作为Key，然后值里面存的是数量 numMap1 = map[4:2 9:1]
	3. 另外一个数组 arr2 也进行遍历，这次以 key 去找数量，找到了者新初始化一个索引，每次找到，把找到的值
放入到新的数组（这里直接用arr2赋值了），然后将该值，对应的数量 -1
*/
func findForMap(arr1, arr2 []int) []int {
	numMap := map[int]int{}
	for _, v := range arr1 {
		numMap[v] += 1
	}

	k := 0

	for _, v := range arr2 {
		if numMap[v] > 0 {
			// 数量减 1
			numMap[v] -= 1
			arr2[k] = v
			k++
		}
	}
	return arr2[0:k]
}

/**
这是遍历有序的数组
*/
func findForSorted(arr1, arr2 []int) []int {
	i, j, k := 0, 0, 0
	sort.Ints(arr1)
	sort.Ints(arr2)
	if i < len(arr1) && j < len(arr2) {
		if arr1[i] > arr2[j] {
			j++
		} else if arr1[i] < arr2[j] {
			i++
		} else {
			// 赋值，还是用arr1去接收，并未去新初始化数组
			arr1[k] = arr1[i]
			i++
			j++
			k++
		}
	}
	return arr1
}
