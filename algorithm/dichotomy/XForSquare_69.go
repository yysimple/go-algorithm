package main

import "fmt"

/**
x的平方根， lc 69
*/
func main() {
	square := xSquare(15)
	fmt.Println("square: ", square)
}

/**
mid := (left+right)>>1 + 1，二分法这里一般加上 1 ，最有一left返回就行了，这不是最完美的二分
这里可以把 +1 操作放到 else中，但是因为最后一次走else的时候，还会加 1 ，所以再返回值的时候 -1 也是可以的 可以试试
*/
func xSquare(target int) int {
	left := 0
	right := target / 2
	for left < right {
		mid := (left+right)>>1 + 1
		if mid > (target / mid) {
			right = mid - 1
		} else {
			left = mid
		}
	}
	return left
}
