package dynamic_planning

/*
给定 M*N 的网格, 有一个机器人从左上角(0,0) 出发, 每一步可以向下走一步或者向右走一步,
问 有多少种不同的方式走到右下角
*/

func MinMoveSteps(arr [][]int) int {
	var res int
	if len(arr) == 0 {
		return res
	}
	if len(arr[0]) == 0 {
		return res
	}
	var row = len(arr)
	var col = len(arr[0])
	var dp = make([][]int, row, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col, col)
	}
	// 计算边界 右边界和 下边界
	dp[0][0] = 1
	for i := 1; i < row; i++ {
		dp[i][0] = 1
	}
	for i := 1; i < col; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if i-1 >= 0 {
				dp[i][j] += dp[i-1][j]
			}
			if j-1 >= 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}
	res = dp[row-1][col-1]
	return res
}
