package main

import "fmt"

/**
24点游戏 lc 679
*/
func main() {
	nums := []int{4, 1, 8, 7}
	point24_3 := judgePoint24_3(nums)
	fmt.Println("ope to 24 =", point24_3)
}

func judgePoint24_1(a, b float64) bool {
	return (a+b < 24+1e-6 && a+b > 24-1e-6) ||
		(a*b < 24+1e-6 && a*b > 24-1e-6) ||
		(a-b < 24+1e-6 && a-b > 24-1e-6) ||
		(b-a < 24+1e-6 && b-a > 24-1e-6) ||
		(a/b < 24+1e-6 && a/b > 24-1e-6) ||
		(b/a < 24+1e-6 && b/a > 24-1e-6)
}

func judgePoint24_2(a, b, c float64) bool {
	return judgePoint24_1(a+b, c) ||
		judgePoint24_1(a-b, c) ||
		judgePoint24_1(a*b, c) ||
		judgePoint24_1(a/b, c) ||
		judgePoint24_1(b-a, c) ||
		judgePoint24_1(b/a, c) ||
		judgePoint24_1(a+c, b) ||
		judgePoint24_1(a-c, b) ||
		judgePoint24_1(a*c, b) ||
		judgePoint24_1(a/c, b) ||
		judgePoint24_1(c-a, b) ||
		judgePoint24_1(c/a, b) ||
		judgePoint24_1(c+b, a) ||
		judgePoint24_1(c-b, a) ||
		judgePoint24_1(c*b, a) ||
		judgePoint24_1(c/b, a) ||
		judgePoint24_1(b-c, a) ||
		judgePoint24_1(b/c, a)
}

func judgePoint24_3(nums []int) bool {
	return judgePoint24_2(float64(nums[0])+float64(nums[1]), float64(nums[2]), float64(nums[3])) ||
		judgePoint24_2(float64(nums[0])-float64(nums[1]), float64(nums[2]), float64(nums[3])) ||
		judgePoint24_2(float64(nums[0])*float64(nums[1]), float64(nums[2]), float64(nums[3])) ||
		judgePoint24_2(float64(nums[0])/float64(nums[1]), float64(nums[2]), float64(nums[3])) ||
		judgePoint24_2(float64(nums[1])-float64(nums[0]), float64(nums[2]), float64(nums[3])) ||
		judgePoint24_2(float64(nums[1])/float64(nums[0]), float64(nums[2]), float64(nums[3])) ||
		judgePoint24_2(float64(nums[0])+float64(nums[2]), float64(nums[1]), float64(nums[3])) ||
		judgePoint24_2(float64(nums[0])-float64(nums[2]), float64(nums[1]), float64(nums[3])) ||
		judgePoint24_2(float64(nums[0])*float64(nums[2]), float64(nums[1]), float64(nums[3])) ||
		judgePoint24_2(float64(nums[0])/float64(nums[2]), float64(nums[1]), float64(nums[3])) ||
		judgePoint24_2(float64(nums[2])-float64(nums[0]), float64(nums[1]), float64(nums[3])) ||
		judgePoint24_2(float64(nums[2])/float64(nums[0]), float64(nums[1]), float64(nums[3])) ||
		judgePoint24_2(float64(nums[0])+float64(nums[3]), float64(nums[2]), float64(nums[1])) ||
		judgePoint24_2(float64(nums[0])-float64(nums[3]), float64(nums[2]), float64(nums[1])) ||
		judgePoint24_2(float64(nums[0])*float64(nums[3]), float64(nums[2]), float64(nums[1])) ||
		judgePoint24_2(float64(nums[0])/float64(nums[3]), float64(nums[2]), float64(nums[1])) ||
		judgePoint24_2(float64(nums[3])-float64(nums[0]), float64(nums[2]), float64(nums[1])) ||
		judgePoint24_2(float64(nums[3])/float64(nums[0]), float64(nums[2]), float64(nums[1])) ||
		judgePoint24_2(float64(nums[2])+float64(nums[3]), float64(nums[0]), float64(nums[1])) ||
		judgePoint24_2(float64(nums[2])-float64(nums[3]), float64(nums[0]), float64(nums[1])) ||
		judgePoint24_2(float64(nums[2])*float64(nums[3]), float64(nums[0]), float64(nums[1])) ||
		judgePoint24_2(float64(nums[2])/float64(nums[3]), float64(nums[0]), float64(nums[1])) ||
		judgePoint24_2(float64(nums[3])-float64(nums[2]), float64(nums[0]), float64(nums[1])) ||
		judgePoint24_2(float64(nums[3])/float64(nums[2]), float64(nums[0]), float64(nums[1])) ||
		judgePoint24_2(float64(nums[1])+float64(nums[2]), float64(nums[0]), float64(nums[3])) ||
		judgePoint24_2(float64(nums[1])-float64(nums[2]), float64(nums[0]), float64(nums[3])) ||
		judgePoint24_2(float64(nums[1])*float64(nums[2]), float64(nums[0]), float64(nums[3])) ||
		judgePoint24_2(float64(nums[1])/float64(nums[2]), float64(nums[0]), float64(nums[3])) ||
		judgePoint24_2(float64(nums[2])-float64(nums[1]), float64(nums[0]), float64(nums[3])) ||
		judgePoint24_2(float64(nums[2])/float64(nums[1]), float64(nums[0]), float64(nums[3])) ||
		judgePoint24_2(float64(nums[1])+float64(nums[3]), float64(nums[2]), float64(nums[0])) ||
		judgePoint24_2(float64(nums[1])-float64(nums[3]), float64(nums[2]), float64(nums[0])) ||
		judgePoint24_2(float64(nums[1])*float64(nums[3]), float64(nums[2]), float64(nums[0])) ||
		judgePoint24_2(float64(nums[1])/float64(nums[3]), float64(nums[2]), float64(nums[0])) ||
		judgePoint24_2(float64(nums[3])-float64(nums[1]), float64(nums[2]), float64(nums[0])) ||
		judgePoint24_2(float64(nums[3])/float64(nums[1]), float64(nums[2]), float64(nums[0]))
}
