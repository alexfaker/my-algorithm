package dynamic_planning

import (
	"math"

	"gitee.com/Anderson/my-algorithm/public"
)

// MaxCutRodProfit 钢筋切割问题的最大收益
// @Description:
// @param prices
// @param n
// @return int
func MaxCutRodProfit(prices []int, n int) int {
	if len(prices) == 0 {
		return 0
	}
	return maxCutRodProfit(prices, n)
}

// 代码就这么简单
func maxCutRodProfit(prices []int, n int) int {
	if n == 0 {
		return 0
	}
	var q = math.MinInt64
	for i := 1; i <= n; i++ {
		q = public.Max(q, prices[i-1]+maxCutRodProfit(prices, n-i))
	}
	return q
}

// MaxCutRodProfit1
// @Description:  动态规划的解法 1 自顶向下解
// @param prices
// @param n
// @return int
func MaxCutRodProfit1(prices []int, n int) int {
	if len(prices) == 0 {
		return 0
	}
	var dp = make([]int, n, n)
	for i := 0; i < n; i++ {
		dp[i] = math.MinInt64
	}
	return maxCutRodProfit1(prices, dp, n)
}

func maxCutRodProfit1(prices, dp []int, n int) int {
	if dp[n-1] >= 0 {
		return dp[n-1]
	}
	var q = math.MinInt64
	if n == 0 {
		q = 0
	}
	for i := 1; i <= n; i++ {
		q = public.Max(q, prices[i-1]+maxCutRodProfit(prices, n-i))
	}
	dp[n-1] = q
	return q
}

// MaxCutRodProfit2
// @Description:  动态规划的解法2 自底向上解
// 时间复杂度和空间复杂度均为 O(logN)
// @param prices
// @param n
// @return int
func MaxCutRodProfit2(prices []int, n int) int {
	if len(prices) == 0 {
		return 0
	}
	var dp = make([]int, n+1, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MinInt64
	}

	for j := 1; j <= n; j++ {
		q := math.MinInt64
		for i := 1; i <= j; i++ {
			q = public.Max(q, prices[i-1]+dp[j-i])
		}
		dp[j] = q
	}
	return dp[n]

}

// MaxCutRodProfit3 动态规划扩展
// @Description:  不止输出最大收益, 也输出最大收益的切割方案
// @param prices
// @param n
// @return int []int
func MaxCutRodProfit3(prices []int, n int) (int, []int) {
	if len(prices) == 0 {
		return 0, nil
	}
	var dp = make([]int, n+1, n+1)
	var res = make([]int, n+1, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MinInt64
	}

	for j := 1; j <= n; j++ {
		q := math.MinInt64
		for i := 1; i <= j; i++ {
			if q < prices[i-1]+dp[j-i] {
				q = prices[i-1] + dp[j-i]
				res[j] = i
			}
		}
		dp[j] = q
	}
	return dp[n], res

}
