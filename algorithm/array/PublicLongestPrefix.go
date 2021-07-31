package main

import (
	"fmt"
	"strings"
)

/**
最长公共前缀， lc：14题
*/
func main() {
	// index = -1
	// index1 := strings.Index("flower", "erflow")
	// index = 0
	// index2 := strings.Index("flower", "flow")
	strs := []string{"abcdef", "abcfrs", "abdfg"}
	api := findByStringApi(strs)
	fmt.Println("longest prefix =", api)
}

func findByStringApi(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	prefix := strs[0]
	for _, str := range strs {
		// 如果 str 从 0号字符开始，包含了 prefix，就会返回 0
		for strings.Index(str, prefix) != 0 {
			if len(prefix) == 0 {
				return ""
			}
			// 一直循环，直到找到最长的
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}
