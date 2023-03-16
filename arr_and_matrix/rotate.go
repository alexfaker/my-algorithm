package arr_and_matrix

func Rotate(matrix [][]int) [][]int {
	var n = len(matrix)
	var res = make([][]int, n, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			tmp := matrix[i][j]
			var x, y = j, i
			if y == 0 {
				y = n - 1
			} else if y == n-1 {
				y = 0
			}
			res[x][y] = tmp
		}
	}
	return res
}

func GameOfLife(board [][]int) [][]int {
	var m = len(board)
	var n = len(board[0])
	var newBoard = make([][]int, m, m)
	for i := 0; i < m; i++ {
		newBoard[i] = make([]int, n, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			liveCnt := getLiveCnt(board, i, j, m, n)
			newBoard[i][j] = board[i][j]
			if board[i][j] == 1 {
				switch liveCnt {
				case 0, 1:
					newBoard[i][j] = 0
					break
				case 2, 3:
					newBoard[i][j] = 1
					break
				default:
					newBoard[i][j] = 0
					break
				}
			} else {
				if liveCnt == 3 {
					newBoard[i][j] = 1
				}
			}
		}
	}
	board = newBoard
	return newBoard
}

func getLiveCnt(board [][]int, i, j, m, n int) int {
	var cnt int
	if i-1 >= 0 {
		cnt += board[i-1][j]
		if j-1 >= 0 {
			cnt += board[i-1][j-1] + board[i][j-1]
		}
		if j+1 < n {
			cnt += board[i-1][j+1] + board[i][j+1]
		}
	}

	if i+1 < m {
		cnt += board[i+1][j]
		if j-1 >= 0 {
			cnt += board[i+1][j-1]
		}
		if j+1 < n {
			cnt += board[i+1][j+1]
		}
	}
	return cnt

}
