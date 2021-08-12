package main

import "fmt"

/**
大数打印：输入 n ，返回n对应位数最大的值，比如 3 就打印 1-999
*/
func main() {
	n := 3
	nums := printBigNum(n)
	fmt.Println("nums =", nums)
}

func printBigNum(n int) []int {
	var res []int
	index := 0
	for 0 < n {
		n--
		/**
		这里没每次 90 结尾 在 + 9最后就可以要得到需要打印的索引数，也就是个数
		*/
		index = index*10 + 9
	}

	for i := 0; i <= index; i++ {
		res = append(res, i)
	}
	return res
}
