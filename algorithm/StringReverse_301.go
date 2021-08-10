package main

import "fmt"

/**
反转字符串，lc 301
*/
func main() {
	str := []string{"h", "e", "l", "l", "o"}
	runes := reverseString(str)
	fmt.Println("rune =", runes)
}

/**
双指针法
*/
func reverseString(s []string) []string {
	left := 0
	right := len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	return s
}
