package dynamic_planning

import "math"

/*
给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。
注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
示例 1:
输入：prices = [3,3,5,0,0,3,1,4]
输出：6
解释：在第 4 天（股票价格 = 0）的时候买入，在第 6 天（股票价格 = 3）的时候卖出，这笔交易所能获得利润 = 3-0 = 3 。
     随后，在第 7 天（股票价格 = 1）的时候买入，在第 8 天 （股票价格 = 4）的时候卖出，这笔交易所能获得利润 = 4-1 = 3 。
链接：https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii
 */

func StockSell(arr []int) int {
	var l = len(arr)
	if l <= 1 {
		return 0
	}
	buy1, sell1 := -arr[0], 0
	buy2, sell2 := -arr[0], 0
	for i := 1; i < l; i++ {
		buy1 = int(math.Max(float64(buy1), float64(-arr[i])))
		sell1 = int(math.Max(float64(buy1+arr[i]), float64(sell1)))
		buy2 = int(math.Max(float64(sell1-arr[i]), float64(buy2)))
		sell2 = int(math.Max(float64(buy2+arr[i]), float64(sell2)))
	}
	if sell2 >= sell1 {
		return sell2
	}
	return sell1
}
