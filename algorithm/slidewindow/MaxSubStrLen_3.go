package main

import "fmt"

/**
无重复字符的最长子串,lc 3
*/
func main() {
	strs := []string{"a", "b", "c", "a", "b", "c", "d", "b", "b"}
	maxStr := findMaxStr(strs)
	fmt.Println("maxStr =", maxStr)
}

func findMaxStr(strs []string) int {
	strArrLen := len(strs)
	maxStrLen := 0
	tmpArr := make([]string, 0)
	for i := 0; i < strArrLen; i++ {
		if !stringsContains(tmpArr, strs[i]) {
			tmpArr = append(tmpArr, strs[i])
		} else {
			maxStrLen = max(maxStrLen, len(tmpArr))
			tmpArr = nil
			i--
		}
	}
	return maxStrLen
}

func stringsContains(array []string, val string) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
