package main

import (
	"fmt"
	"math"
)

/**
灯泡开关，lc 319
*/
func main() {
	n := 12
	num := lightNum(n)
	fmt.Println("第", n, "轮之后,", "剩余的灯数：", int(num))
	lightNum2(n)
}

/**
可以根据题目意思，然后自己推导，最后会得出一个结论：开根号的数值取整就是最后剩余的灯数
*/
func lightNum(n int) float64 {
	return math.Sqrt(float64(n))
}

/**
这里的话，就是第一轮全开，第二轮每两个切换，第三轮，每三个切换...以此推导，即使可以得这个公式 == (j+1)%i == 0
*/
func lightNum2(n int) {
	var light = make([]int, n)
	for i := 0; i < n; i++ {
		light = append(light[:i], append([]int{-1}, light[i:]...)...)
	}
	light = append(light[0:n])
	fmt.Println("初始化灯泡：", light)

	for i := 1; i <= n; i++ {
		for j := 0; j < n; j++ {
			if (j+1)%i == 0 {
				light[j] = light[j] * -1
			}
		}
		fmt.Println("第", i, "轮，灯泡：", light)
	}
}
