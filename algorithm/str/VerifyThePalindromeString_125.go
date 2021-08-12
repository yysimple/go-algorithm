package main

import (
	"fmt"
	"strings"
)

/**
验证回文串，lc 125
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
*/
func main() {
	str := "A man, a plan, a canal: Panama"
	palindrome := checkPalindrome(str)
	fmt.Println("yes ?? = ", palindrome)
	b := isPalindrome(str)
	fmt.Println("b =", b)
	fmt.Println("check =", str[15] == str[18])
}

/**
这里是去掉不符合的字符，然后双指针法进行遍历
*/
func checkPalindrome(str string) bool {
	var onlyStr string
	for i := 0; i < len(str); i++ {
		if isCharOrNum(str[i]) {
			onlyStr += string(str[i])
		}
	}

	strLen := len(onlyStr)
	onlyStr = strings.ToLower(onlyStr)
	fmt.Println("onlyStr =", onlyStr)
	for i := 0; i < strLen/2; i++ {
		if onlyStr[i] != onlyStr[strLen-1-i] {
			return false
		}
	}
	return true
}

/**
这种就是两边同时移动，碰到不是字符的就跳过
*/
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isCharOrNum(s[left]) {
			left++
		}
		for left < right && !isCharOrNum(s[right]) {
			right--
		}
		// 这里判断
		if left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
	}
	return true
}

/**
校验是否是字符或者数字
*/
func isCharOrNum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
