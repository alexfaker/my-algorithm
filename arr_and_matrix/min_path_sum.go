package arr_and_matrix

import (
	"math"
)

/*
求解矩阵的最小路径和问题
给定一个矩阵,从左上角开始, 每次只能向下或者向右走,最后到达右下角的位置
路径上所有的数字累加起来 就是路径和, 返回所有的路径和中最小的路径和
*/

/*
解法 1: 经典的动态规划 生成一个 M*N 的动态规划数组, 记录走到每一步的最小路径和
时间复杂度 O(M*N) 空间复杂度 O(M*N)

*/

func MinPathSum(m [][]int) int {
	var res int
	if len(m) <= 0 {
		return res
	}
	var dp = make([][]int, len(m), len(m))
	for i := 0; i < len(m); i++ {
		dp[i] = make([]int, len(m[0]), len(m[0]))
	}

	//计算右边界和左边界的值
	for i := 0; i < len(m[0]); i++ {
		if i == 0 {
			dp[0][i] = m[0][i]
		} else {
			dp[0][i] = dp[0][i-1] + m[0][i]
		}
	}

	for i := 0; i < len(m); i++ {
		if i == 0 {
			dp[i][0] = m[i][0]
		} else {
			dp[i][0] = dp[i-1][0] + m[i][0]
		}
	}

	//计算m 每个值的最小路径和
	for i := 1; i < len(m); i++ {
		for j := 1; j < len(m[0]); j++ {
			dp[i][j] = int(math.Min(float64(dp[i-1][j]), float64(dp[i][j-1]))) + m[i][j]
		}
	}

	res = dp[len(m)-1][len(m[0])-1]

	return res
}

/*
解法 2 ,数据空间压缩后的动态规划
*/

func MinPathSum2(m [][]int) int {
	if len(m) == 0 || len(m[0]) == 0 {
		return 0
	}
	var more = int(math.Max(float64(len(m)), float64(len(m[0]))))
	var less = int(math.Min(float64(len(m)), float64(len(m[0]))))
	var rowMore = more == len(m)
	var arr = make([]int, less, less)
	arr[0] = m[0][0]
	for i := 1; i < less; i++ {
		if rowMore {
			arr[i] = arr[i-1] + m[0][i]
		} else {
			arr[i] = arr[i-1] + m[i][0]
		}
	}

	for i := 1; i < more; i++ {
		arr[0] += func() int {
			if rowMore {
				return m[i][0]
			}
			return m[0][i]
		}()
		for j := 1; j < less; j++ {
			arr[j] = int(math.Min(float64(arr[j-1]), float64(arr[j]))) + func() int {
				if rowMore {
					return m[i][j]
				}
				return m[j][i]
			}()
		}
	}
	return arr[less-1]
}
