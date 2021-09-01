package main

import "fmt"

/**
给定一个整数，编写一个函数来判断它是否是 2 的幂次方。
*/
func main() {
	two := isPowerOfTwo(16)
	fmt.Println("is power of 2: ", two)
}

func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}
