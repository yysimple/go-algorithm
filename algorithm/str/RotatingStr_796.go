package main

import (
	"fmt"
	"strings"
)

/**
旋转字符串，lc 796
*/
func main() {
	str1 := "abcde"
	str2 := "cdeab"
	str := rotatingStr(str1, str2)
	fmt.Println("str =", str)

}

/**
旋转数组
*/
func rotatingStr(fatherStr string, sonStr string) bool {
	fatherStr = fatherStr + fatherStr
	return strings.Contains(fatherStr, sonStr)
}
