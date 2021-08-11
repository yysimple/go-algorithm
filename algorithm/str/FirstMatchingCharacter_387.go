package main

import "fmt"

/**
字符串中的第一个唯一字符，lc 387
*/
func main() {
	str := "helloworldhehe"
	printChar(str)
	char := firstUniqChar(str)
	byMap := firstUniqCharByMap(str)
	fmt.Println("char =", char)
	fmt.Println("byMap =", byMap)
}

/**
使用 [26]int 记录对应字符出现在字符串中的索引
*/
func firstUniqChar(s string) int {
	var arr [26]int
	// 遍历索引，然后将相同的值放入,如果存在就覆盖；
	for i, k := range s {
		arr[k-'a'] = i
	}
	// 再次遍历，这里要是遍历的时候，i索引在数组中找不到，则证明其不是唯一，则将值置为 -1
	for i, k := range s {
		if i == arr[k-'a'] {
			return i
		} else {
			arr[k-'a'] = -1
		}
	}
	return -1
}

/**
这里是使用map完成的， 如果字符只出现一次就记录她的索引，反之将其记对应的值改为 -1，再次遍历的时候，只要对应的 v， 在map中获取的值 == i就行
*/
func firstUniqCharByMap(s string) int {
	uniqCharMap := make(map[int]int)
	for index, value := range s {
		if uniqCharMap[int(value)] != 0 {
			uniqCharMap[int(value)] = -1
		}
		uniqCharMap[int(value)] = index
	}

	for index, value := range s {
		if uniqCharMap[int(value)] == index {
			return index
		}
	}
	return -1
}

func printChar(s string) {
	var arr [26]int
	for i, k := range s {
		arr[k-'a'] = i
	}
	fmt.Println("arr =", arr)
}
