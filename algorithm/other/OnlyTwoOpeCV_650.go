package main

import "fmt"

/**
只有两个键的键盘, lc 650
*/
func main() {
	steps := minSteps(30)
	fmt.Println("steps =", steps)
}

/**
这里其实求的就是素素相加就拆分一下30 = 2 * 3 * 5，都是素数，然后在相加 2 + 3 + 5 = 10
*/
func minSteps(n int) int {
	res := 0
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			res += i
			n /= i
		}
	}
	return res
}
