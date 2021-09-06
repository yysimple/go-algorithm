package main

import "fmt"

/**
返回二进制中1的个数,lc 191
*/
func main() {
	weight := hammingWeight(31)
	fmt.Println("nums =", weight)
}

/**
这里比较简单，就是利用 & 的特性，相同则位1，这里只需要掩码 0001 的 1 一直左移就行，然后 & 运算之后为1则 ++
*/
func hammingWeight(target int) int {
	res := 0
	mask := 1
	for i := 0; i < 32; i++ {
		if (target & mask) != 0 {
			res++
		}
		mask = mask << 1
	}
	return res
}
