package main

import "fmt"

/**
只有两个键的键盘, lc 650

*/
func main() {
	steps := minSteps(10)
	fmt.Println("steps =", steps)
}

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
