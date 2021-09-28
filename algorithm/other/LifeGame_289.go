package main

import "fmt"

/**
生命游戏（289）
*/
func main() {
	board := [][]int{[]int{0, 1, 0}, []int{0, 0, 1}, []int{1, 1, 1}, []int{0, 0, 0}}
	liftGame(board)
	fmt.Println("board: ", board)
}

func liftGame(board [][]int) {
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			liveNeighbors := 0
			// 上面
			if i > 0 {
				liveNeighbors += sanMuExp(board[i-1][j] == 1, board[i-1][j] == 2)
			}
			// 判断左边
			if j > 0 {
				liveNeighbors += sanMuExp(board[i][j-1] == 1, board[i][j-1] == 2)
			}
			// 判断下边
			if i < m-1 {
				liveNeighbors += sanMuExp(board[i+1][j] == 1, board[i+1][j] == 2)
			}
			// 判断右边
			if j < n-1 {
				liveNeighbors += sanMuExp(board[i][j+1] == 1, board[i][j+1] == 2)
			}
			// 判断左上角
			if i > 0 && j > 0 {
				liveNeighbors += sanMuExp(board[i-1][j-1] == 1, board[i-1][j-1] == 2)
			}
			//判断右下角
			if i < m-1 && j < n-1 {
				liveNeighbors += sanMuExp(board[i+1][j+1] == 1, board[i+1][j+1] == 2)
			}
			// 判断右上角
			if i > 0 && j < n-1 {
				liveNeighbors += sanMuExp(board[i-1][j+1] == 1, board[i-1][j+1] == 2)
			}
			// 判断左下角
			if i < m-1 && j > 0 {
				liveNeighbors += sanMuExp(board[i+1][j-1] == 1, board[i+1][j-1] == 2)
			}
			// 根据周边存活数量更新当前点，结果是 0 和 1 的情况不用更新
			if board[i][j] == 0 && liveNeighbors == 3 {
				board[i][j] = 3
			} else if board[i][j] == 1 {
				if liveNeighbors < 2 || liveNeighbors > 3 {
					board[i][j] = 2
				}
			}
		}
	}
	// 过渡状态 -> 真实状态
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = board[i][j] % 2
		}
	}
}

func sanMuExp(var1, var2 bool) int {
	if var1 || var2 {
		return 1
	} else {
		return 0
	}
}
