package main

import "fmt"

/**
螺旋矩阵Ⅰ,lc 54
*/
func main() {
	nums := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	order := spiralOrder(nums)
	fmt.Println("order: ", order)
}

func spiralOrder(matrix [][]int) []int {
	var result []int
	if len(matrix) == 0 {
		return result
	}
	// 这里对应数组的最左边、最右边、最上边、最下边
	left, right, up, down := 0, len(matrix[0])-1, 0, len(matrix)-1
	var x, y int
	for avoid(up, down, left, right) {
		// 左到右
		for y = left; y <= right; y++ {
			result = append(result, matrix[x][y])
		}
		// 最后一个y多++了一次
		y--
		// 行 往下移一位
		up++
		// 上到下
		for x = up; x <= down; x++ {
			result = append(result, matrix[x][y])
		}
		x--
		// 最后一列输出完毕，往左移一列
		right--
		for y = right; y >= left; y-- {
			result = append(result, matrix[x][y])
		}
		// 多减了 1
		y++
		// 最后一列输出完毕，往左移一列
		down--
		for x = down; x >= up; x-- {
			result = append(result, matrix[x][y])
		}
		x++
		left++
	}
	return result
}

func avoid(up, down, left, right int) bool {
	return up <= down && left <= right
}
