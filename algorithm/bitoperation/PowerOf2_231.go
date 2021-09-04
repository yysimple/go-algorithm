package main

import "fmt"

/**
给定一个整数，编写一个函数来判断它是否是 2 的幂次方。
*/
func main() {
	two := isPowerOfTwo(64)
	fmt.Println("is power of 2: ", two)
}

/**
& 表示的是:只要不为1则为零：
我们知道，2的幂次方一般转成二进制就是：1000 0000 = 128 / 0100 0000 = 64 / 0010 0000 = 32 / 0001 0000 = 16
可以从此种我们推断出，我们只需要找到他的 -1 数，比如 16 - 1 = 15 / 0000 1111，那么做 & 运算的时候就会等于 0 ，然后在递归求解就行
*/
func isPowerOfTwo(n int) bool {
	if n&(n-1) == 0 {
		return true
	}
	return false
}
