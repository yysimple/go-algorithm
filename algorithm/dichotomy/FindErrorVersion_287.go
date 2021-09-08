package main

import "fmt"

/**
第一个错误的版本 lc，287
*/
func main() {
	version := findFirstErrorVersion(7)
	fmt.Println("version: ", version)
}

/**
输出总共有多少版本。假设有一个知道版本是失败的方法，这里用isBadVersion来模拟
然后找到对应的第一个失败的版本
*/
func findFirstErrorVersion(version int) int {
	left := 1
	right := version
	for left < right {
		// 这里是从版本 1 开始的，所以把left加上
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

/**
这里指定，第四个版本是第一个错误版本
*/
func isBadVersion(targetVersion int) bool {
	if targetVersion <= 3 {
		return false
	}
	return true
}
