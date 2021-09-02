package main

import "fmt"

/**
返回二进制中1的个数,lc 191
*/
func main() {
	weight := hammingWeight(31)
	fmt.Println("nums =", weight)
}

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
