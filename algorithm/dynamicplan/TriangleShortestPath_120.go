package main

import "fmt"

/**
三角型最小路径，lc 120
给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
*/
func main() {
	triangle := buildTriangle()
	printTwoArr(triangle)
	total := minimumTotal(triangle)
	fmt.Println("total =", total)

}

func minimumTotal(triangle [][]int) int {
	l := len(triangle)
	if l < 1 {
		return 0
	}
	if l == 1 {
		return triangle[0][0]
	}
	result := 1<<31 - 1
	// 初始化前两行数据；
	triangle[0][0] = triangle[0][0]
	triangle[1][1] = triangle[1][1] + triangle[0][0]
	triangle[1][0] = triangle[1][0] + triangle[0][0]
	// 从第三行开始遍历，这里是开始累积每一条路径的值，最后取最小的
	for i := 2; i < l; i++ {
		// 每行开始遍历
		for j := 0; j < len(triangle[i]); j++ {
			// j=0，则累加上一行的0列元素
			if j == 0 {
				triangle[i][j] = triangle[i-1][j] + triangle[i][j]
				// 这里如果是最右边的元素，则只累加上一行 上一列的值，但是这里我初始化数组会出现很多0，所以需要处理一下
			} else if j == (getArrLen(triangle[i]) - 1) {
				triangle[i][j] = triangle[i-1][j-1] + triangle[i][j]
				// 这里就是去上一行当前列/上一行上一列中的较小值即可
			} else {
				triangle[i][j] = min(triangle[i-1][j-1], triangle[i-1][j]) + triangle[i][j]
			}
		}
	}
	// 这里取出最后一行的值，然后取最小的，就是最小路劲
	for _, k := range triangle[l-1] {
		result = min(result, k)
	}
	return result
}

func getArrLen(arr []int) int {
	flag := 0
	for _, k := range arr {
		if k > 0 {
			flag++
		}
	}
	return flag
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func buildTriangle() [][]int {
	x := 4
	y := 4

	triangle := make([][]int, x)
	for i := range triangle {
		triangle[i] = make([]int, y)
	}

	triangle[0][0] = 2
	triangle[1][0] = 3
	triangle[1][1] = 4
	triangle[2][0] = 6
	triangle[2][1] = 5
	triangle[2][2] = 7
	triangle[3][0] = 4
	triangle[3][1] = 1
	triangle[3][2] = 8
	triangle[3][3] = 3
	return triangle
}

func printTwoArr(arr [][]int) {
	for i, arrOne := range arr {
		for j := range arrOne {
			fmt.Print("[", arr[i][j], "] ")
		}
		fmt.Println("")
	}
}
