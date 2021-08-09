package main

import "fmt"

/**
最小路劲，lc 64
*/
func main() {
	grid := buildGrid()
	printTwoArr(grid)
	sum := minPathSum(grid)
	fmt.Println("sum =", sum)

}

/**
这个解法跟上一个三角形的差不多，也是一直往下累加
*/
func minPathSum(grid [][]int) int {
	arrLen := len(grid)
	if arrLen < 1 {
		return 0
	}
	for i := 0; i < arrLen; i++ {
		for j := 0; j < len(grid[i]); j++ {
			if j == 0 && i != 0 {
				grid[i][j] = grid[i-1][j] + grid[i][j]
			} else if j != 0 && i == 0 {
				grid[i][j] = grid[i][j-1] + grid[i][j]
			} else if i != 0 && j != 0 {
				grid[i][j] = min(grid[i-1][j], grid[i][j-1]) + grid[i][j]
			}
		}
	}
	return grid[arrLen-1][len(grid[arrLen-1])-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func buildGrid() [][]int {
	x := 3
	y := 3

	buildGrid := make([][]int, x)
	for i := range buildGrid {
		buildGrid[i] = make([]int, y)
	}

	buildGrid[0][0] = 1
	buildGrid[0][1] = 3
	buildGrid[0][2] = 1
	buildGrid[1][0] = 1
	buildGrid[1][1] = 5
	buildGrid[1][2] = 1
	buildGrid[2][0] = 4
	buildGrid[2][1] = 2
	buildGrid[2][2] = 1
	return buildGrid
}

func printTwoArr(arr [][]int) {
	for i, arrOne := range arr {
		for j := range arrOne {
			fmt.Print("[", arr[i][j], "] ")
		}
		fmt.Println("")
	}
}
