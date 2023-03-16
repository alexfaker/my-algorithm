package dynamic_planning

import (
	"math"

	"gitee.com/Anderson/my-algorithm/public"
)

/*
最大子数组问题, 递归解法
给定一个数组. 值有正数和负数,
求该数组内连续子数组的加和记为 N
求 N 的最大值

子问题分解:
左半部分数组的最大子数组和
右半部分的最大子数组和
跨越 mid 的最大子数组和

原问题的解是以上三种情况中的最大值

时间复杂度 O(NlogN)	空间复杂度 O(1)
*/

func MaxSubArr(arr []int) int {
	var l = len(arr)
	if l == 0 {
		return 0
	}
	return maxSubArr(arr, 0, l-1)

}

func maxSubArr(arr []int, left, right int) int {
	if left == right {
		return arr[left]
	}
	var mid = (left + right) / 2
	// 向左边要数据
	leftSum := maxSubArr(arr, left, mid)
	// 向右边要数据
	rightSum := maxSubArr(arr, mid+1, right)
	// 计算跨越终点的值
	midSum := getMidSum(arr, left, mid, right)
	// 取最大值
	return public.Max(public.Max(leftSum, rightSum), midSum)

}

/*
跨越边界的子数组最大和
 */
func getMidSum(arr []int, left, mid, right int) int {
	var leftSum, rightSum = math.MinInt64, math.MinInt64
	var sum = 0
	// mid 到左边界的最大值
	for i := mid; i >= left; i-- {
		sum += arr[i]
		if sum > leftSum {
			leftSum = sum
		}
	}
	sum = 0
	// mid 到右边界的最大值
	for i := mid + 1; i <= right; i++ {
		sum += arr[i]
		if sum > rightSum {
			rightSum = sum
		}
	}
	return leftSum + rightSum
}

/*
动态规划版本的实现
时间复杂度 O(N)	空间复杂度 O(N)
递归实现的问题在于有些子问题会重复计算
而动态规划 保存了子问题的解, 从而避免了子问题的重复计算
*/

func MaxSubArr1(arr []int) int {
	var l = len(arr)
	if l == 0 {
		return 0
	}
	var dp = make([]int, l, l)
	dp[0] = arr[0]
	for i := 1; i < l; i++ {
		if dp[i-1] > 0 {
			dp[i] = arr[i] + dp[i-1]
		} else {
			dp[i] = arr[i]
		}
	}

	var maxSum = math.MinInt64
	for _, v := range dp {
		if v > maxSum {
			maxSum = v
		}
	}
	return maxSum

}