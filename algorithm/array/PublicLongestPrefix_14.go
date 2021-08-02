package main

import (
	"fmt"
	"strings"
	"time"
)

/**
最长公共前缀， lc：14题

总结：
	- 这使用的是strings的一个api + 两层轮询，去完成的，拿出多个字符串中的第一个当为对比对象即 prefix := strs[0],
然后遍历所有字符串，二层遍历以 strings.Index(str, prefix) != 0 作为结束条件（可以看看代码里的例子），然后一直截取
prefix[0:len-1],直到最后找到，返回prefix
*/
func main() {
	start := time.Now().UnixMicro()
	fmt.Println("start ===> ", start)
	// index = -1
	// index1 := strings.Index("flower", "erflow")
	// index = 0
	// index2 := strings.Index("flower", "flow")
	strs := []string{"abcdef", "abcfrs", "abdfg"}
	api := findByStringApi(strs)
	fmt.Println("longest prefix =", api)
	// time.Sleep(1)

	end := time.Now().UnixMicro()
	fmt.Println("end ===> ", end)
	total := end - start
	fmt.Println("total time ===> ", total)
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
