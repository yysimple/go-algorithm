package main

import "fmt"

/**
实现 strStr()，lc 28
*/
func main() {
	source := []string{"H", "e", "r", "e", "-", "i", "s", "-", "a", "-", "-", "l", "i", "t", "t", "l", "e", "-", "H", "a", "o"}
	fmt.Println("source =", source)
	target := []string{"l", "i", "t", "t", "l", "e"}
	arr := strStr(source, target)
	fmt.Println("arr =", arr)
	str1 := "Here is a  little Hao"
	str2 := "little"
	str := strStr01(str1, str2)
	fmt.Println("strIndex =", str)
}

/**
数组方式是我不会操作字符串,写了一遍
*/
func strStr(source []string, target []string) int {
	if len(source) < len(target) {
		return -1
	}
	sourceIndex := 0
	targetIndex := 0
	for targetIndex < len(target) {
		if sourceIndex > len(source)-1 {
			return -1
		}
		// 如果元素都相同，则两者继续对比下一个元素
		if source[sourceIndex] == target[targetIndex] {
			sourceIndex++
			targetIndex++
		} else {
			//  len(target) - targetIndex：这个是因为，该target数组也移动了，所以sourceIndex应该移动 target数组长度剩余的部分
			sourceIndex = sourceIndex - targetIndex + len(target)
			if sourceIndex < len(source) {
				step := lastFindIndex(target, source[sourceIndex])
				if step == -1 {
					// 记录此时 source数组遍历的索引位置
					sourceIndex += 1
				} else {
					sourceIndex -= len(target) - step
				}
				// 没找到则从第一个元素开始
				targetIndex = 0
			} else {
				return -1
			}
		}
	}
	return sourceIndex - targetIndex
}

func strStr01(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) == 0 {
		return -1
	}
	if len(needle) > len(haystack) {
		return -1
	}
	// sourceIndex = haystack字符串上的索引， targetIndex = needle字符串上的索引
	sourceIndex, targetIndex := 0, 0
	for targetIndex < len(needle) {
		if sourceIndex >= len(haystack) {
			return -1
		}
		if haystack[sourceIndex] == needle[targetIndex] {
			sourceIndex++
			targetIndex++
		} else {
			sourceIndex = sourceIndex + len(needle) - targetIndex
			if sourceIndex >= len(haystack) {
				return -1
			}
			targetIndex = 0
			tmpIndex := -1
			// 如果 haystack[i] 存在于 needle中 就回退
			// 从后向前匹配
			for k := len(needle) - 1; k >= 0; k-- {
				if needle[k] == haystack[sourceIndex] {
					tmpIndex = len(needle) - k
					break
				}
			}

			if tmpIndex != -1 {
				sourceIndex -= len(needle) - tmpIndex
				targetIndex = 0
			}
		}
	}
	return sourceIndex - targetIndex
}

func lastFindIndex(arr []string, target string) int {
	step := -1
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == target {
			step = len(arr) - i
			break
		}
	}
	return step
}
