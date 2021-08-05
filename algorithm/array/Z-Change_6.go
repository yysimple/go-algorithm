package main

import (
	"fmt"
	"strings"
)

/**
Z字变形，lc 6
总结：这里要悟出一个公式，那就是 2n - 2就是每次循环的一个周期
*/
func main() {

	s := "dasfasfsa"
	// 可以将字符串转成 unicode码 b = [100 97 115 102 97 115 102 115 97]
	b := []rune(s)
	// 这里可以转回成string
	s2 := string(b[2])
	fmt.Println("s2 =", s2)
	fmt.Println("---------------------------")

	s = "LEETCODEISHIRING"
	rows := 4
	change := zChange(s, rows)
	fmt.Println("change =", change)

}

func zChange(s string, rowNum int) string {
	if rowNum == 1 {
		return s
	}

	chars := []rune(s)
	resNums := make([]string, rowNum)
	circle := 2*rowNum - 2
	strLenght := len(chars)
	for i := 0; i < strLenght; i++ {
		index := i % circle
		if index < rowNum {
			resNums[index] += string(chars[i])
		} else {
			resNums[circle-index] += string(chars[i])
		}
	}
	return strings.Join(resNums, "")
}
