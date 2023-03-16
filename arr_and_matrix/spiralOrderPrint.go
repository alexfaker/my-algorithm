package arr_and_matrix

import "fmt"

// 转圈打印矩阵

// PrintMatrixSpiralOrder 思路, 由左上, 右下两个端点确定子矩阵, 完成子矩阵的打印,
// 再缩小端点, 完成内层子矩阵的打印, 直到端点重合
func PrintMatrixSpiralOrder(matrix [][]int) {
	var m = len(matrix[0])
	var n = len(matrix)
	var tr, tc = 0, 0
	var dr, dc = m - 1, n - 1
	for tr <= dr && tc <= dc {
		printEdge(matrix, tr, tc, dr, dc)
		tr++
		tc++
		dr--
		dc--
	}
}

func printEdge(matrix [][]int, tr, tc, dr, dc int) {
	if tr == dr {
		for i := 0; i <= dc; i++ {
			fmt.Printf("%d ", matrix[tr][i])
		}
	} else if tc == dc {
		for i := 0; i <= dr; i++ {
			fmt.Printf("%d ", matrix[i][tc])
		}
	} else {
		curC := tc
		curR := tr
		for curC != dc {
			fmt.Printf("%d ", matrix[tr][curC])
			curC++
		}
		for curR != dr {
			fmt.Printf("%d ", matrix[curR][dc])
			curR++
		}
		for curC != tc {
			fmt.Printf("%d ", matrix[dr][curC])
			curC--
		}
		for curR != tr {
			fmt.Printf("%d ", matrix[curR][tc])
			curR--
		}

	}
}
