package main

import "fmt"

/**
求 1 2 ... n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。
*/
func main() {
	nums := sumNums(3)
	fmt.Println("nums =", nums)
}

func plus(a *int, b int) bool {
	*a += b
	return true
}

/**
这里的思想就是递归往下查，有点类似 斐波那契 求解过程，这是用的是 && 且的求解方式
*/
func sumNums(n int) int {
	_ = n > 0 && plus(&n, sumNums(n-1))
	return n
}
