package main

import "fmt"

/**
Nim 游戏
*/
func main() {
	nim := canWinNim(9)
	fmt.Println("nim: ", nim)
}

/**
这其实是个脑筋急转弯的题目，4的倍数先拿的那个人就输了
*/
func canWinNim(n int) bool {
	return n%4 != 0
}
