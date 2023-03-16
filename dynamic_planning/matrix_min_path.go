package dynamic_planning

import "math"

/*
求解矩阵的最小路径和
时间复杂度 O(m*n), 空间复杂度O(n*m)
*/

func GetMinPathSum(arr [][]int) int {
	var res = -1
	var row = len(arr)
	if row == 0 || len(arr[0]) == 0 {
		return res
	}
	var col = len(arr[0])
	// 初始化 dp 矩阵
	var dp = make([][]int, row, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col, col)
	}
	// 初始化边界条件
	dp[0][0] = arr[0][0]
	for i := 1; i < row; i++ {
		dp[i][0] = dp[i-1][0] + arr[i][0]
	}
	for i := 1; i < col; i++ {
		dp[0][i] = dp[0][i-1] + arr[0][i]
	}

	// 子问题求解
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			dp[i][j] = arr[i][j] + int(math.Min(float64(dp[i-1][j]), float64(dp[i][j-1])))
		}
	}
	res = dp[row-1][col-1]
	return res
}

/*
  压缩版, 不适用矩阵保存动态规划结果, 使用一个数组
  时间复杂度 O(m*n), 空间复杂度O(min{n,m})
*/

func GetMinPathSum1(arr [][]int) int {
	var res = -1
	var row = len(arr)
	if row == 0 || len(arr[0]) == 0 {
		return res
	}
	var col = len(arr[0])

	var dpLen int
	var isRow bool
	if row <= col {
		dpLen = row
		isRow = true
	} else {
		dpLen = col
		isRow = false
	}

	// 初始化 dp 矩阵
	var dp = make([]int, dpLen, dpLen)

	// 初始化边界条件
	dp[0] = arr[0][0]
	if isRow {
		for i := 1; i < row; i++ {
			dp[i] = dp[i-1] + arr[i][0]
		}
	} else {
		for i := 1; i < col; i++ {
			dp[i] = dp[i-1] + arr[0][i]
		}
	}

	// 子问题求
	if isRow {
		for i := 1; i < col; i++ {
			for j := 0; j < row; j++ {
				if j == 0 {
					dp[j] = dp[0] + arr[0][i]
				} else {
					dp[j] = arr[j][i] + int(math.Min(float64(dp[j]), float64(dp[j-1])))
				}
			}
		}
		res = dp[row-1]
	} else {
		for i := 1; i < row; i++ {
			for j := 0; j < col; j++ {
				if j == 0 {
					dp[j] = dp[0] + arr[i][0]
				} else {
					dp[j] = arr[i][j] + int(math.Min(float64(dp[j]), float64(dp[j-1])))
				}
			}
		}
		res = dp[col-1]
	}
	return res
}

func GetMinPathSum2(arr [][]int) int {
	var raw = len(arr)
	if raw == 0 {
		return 0
	}
	var col = len(arr[0])

	var dp = make([]int, col, col)

	// 初始化 dp 数组
	dp[0] = arr[0][0]
	for i := 1; i < col; i++ {
		dp[i] = arr[0][i] + dp[i-1]
	}

	for i := 1; i < raw; i++ {

		for j := 0; j < col; j++ {
			if j == 0 {
				dp[j] = dp[j] + arr[i][j]
			} else {
				dp[j] = arr[i][j] + int(math.Min(float64(dp[j-1]), float64(dp[j])))
			}
		}
	}
	return dp[col-1]
}
