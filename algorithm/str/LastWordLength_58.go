package main

import "fmt"

/**
最后一个单词的长度，lc 58
*/
func main() {
	str := "Hello World "
	wordLen := lastWordLen(str)
	fmt.Println("wordLen =", wordLen)
}

/**
这里第一个字符不能为空
*/
func lastWordLen(str string) int {
	if str == "" || len(str) == 0 {
		return 0
	}
	// 从最后一个一个元素开始遍历，碰到空格跳出循环，长度 -1
	tail := len(str) - 1
	for tail >= 0 && str[tail] == ' ' {
		tail--
	}

	// 赋值给 下一个遍历函数
	head := tail
	// 直到碰到下一个空格之前，否则 --
	for head >= 0 && str[head] != ' ' {
		head--
	}
	// 这里 head 直到 World 的前一个空格位置 head = 5的时候，就再次遇到空格
	return tail - head

}
