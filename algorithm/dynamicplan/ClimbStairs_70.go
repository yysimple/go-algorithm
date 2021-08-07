package main

import "fmt"

/**
爬楼梯问题，lc 70，就是斐波那契问题
*/
func main() {
	way := sumStepWay(6)
	fmt.Println("way =", way)
}

func sumStepWay(step int) int {
	if step == 1 {
		return 1
	}
	if step == 2 {
		return 2
	}
	pre, current := 1, 2
	next := pre + current
	for i := 3; i < step; i++ {
		pre, current = current, next
		next = pre + current
	}
	return next
}
