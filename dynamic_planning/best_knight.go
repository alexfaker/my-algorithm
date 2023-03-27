package dynamic_planning

import "math"

/*
最佳骑士问题
给定一个二维数组 map , 含义是一张地图, 例如
-2 -3 3
-5 -10 1
0 30 -5
游戏规则如下:
* 骑士从左上角出发, 每次只能向右或者向下走, 最后走到右下角拯救公主
* 地图的每个位置的值代表骑士 要遭遇的事情. 负数, 说明此处有怪兽, 打怪兽需要消耗对应的血量,
	非负数, 说明此处有血瓶, 可以恢复骑士对应的血量
* 骑士从左上角走到右下角的过程中, 任何时候血量不能低于 1 ,
问: 为了保证骑士能够顺利拯救公主, 骑士的初始血量最低是多少? 返回最低的初始血量

*/

func MinHp(arr [][]int) int {
	var row = len(arr)
	if row == 0 || len(arr[0]) == 0 {
		return 1
	}
	var col = len(arr[0])
	var dp = make([][]int, row, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col, col)
	}
	if arr[row-1][col-1] > 0 {
		dp[row-1][col-1] = 1
	} else {
		dp[row-1][col-1] = 1 - arr[row-1][col-1]
	}

	for i := row - 2; i >= 0; i-- {
		dp[i][col-1] = int(math.Max(float64(dp[i+1][col-1]-arr[i][col-1]), float64(1)))
	}
	for j := col - 2; j >= 0; j-- {
		dp[row-1][j] = int(math.Max(float64(dp[row-1][j+1]-arr[row-1][j]), float64(1)))
	}
	for i := row - 2; i >= 0; i-- {
		for j := col - 2; j >= 0; j-- {
			right := int(math.Max(float64(dp[i][j+1]-arr[i][j]), float64(1)))
			down := int(math.Max(float64(dp[i+1][j]-arr[i][j]), float64(1)))
			dp[i][j] = int(math.Min(float64(right), float64(down)))
		}
	}

	return dp[0][0]
}
