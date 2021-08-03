package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
加一，lc 66
*/
func main() {
	arr := []int{3, 2, 2, 3}
	ints := plusOneChange(arr)
	fmt.Println("ints =", ints)

	fmt.Println("---------------------------")

	arr1 := []int{3, 2, 2, 3}
	one := plusOne(arr1)
	fmt.Println("one =", one)

}

func plusOne(arr []int) []int {
	var result []int
	addon := 0
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i] += addon
		addon = 0
		if arr[i] == 10 {
			addon = 1
			arr[i] = arr[i] % 10
		}
	}
	if addon == 1 {
		result = make([]int, 1)
		result[0] = 1
		result = append(result, arr...)
	} else {
		result = arr
	}
	return result
}

/**
这种方式就是将 []int ->  string -> int + 1 -> string -> []int
*/
func plusOneChange(arr []int) []int {
	str := ""
	for i := 0; i < len(arr); i++ {
		str = fmt.Sprintf("%s%d", str, arr[i])
	}
	// 字符串转int
	num, _ := strconv.Atoi(str)
	num = num + 1
	// int转字符串
	addAfterString := strconv.Itoa(num)
	return stringToIntArray(addAfterString)
}

/**
string 转换成 int数组
*/
func stringToIntArray(str string) []int {
	split := strings.Split(str, "")
	numArr := make([]int, len(split))
	for i := 0; i < len(split); i++ {
		numArr[i], _ = strconv.Atoi(split[i])
	}
	return numArr
}
