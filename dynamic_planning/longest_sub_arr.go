package dynamic_planning

import "math"

// 获取一个数组的最长递增子序列
// 第一种实现方式, 动态规划
// 先计算 以arr 每个元素结尾的最长递增子序列长度
//  然后再获取最长长度的 序列
// dp[i] 表示以 arr[i] 结尾的最长递增子序列长度

func GetLongestSubArr(arr []int) []int {
	var res []int
	if len(arr) == 0 {
		return res
	}
	dp := getDp1(arr)
	return genLIS(arr, dp)
}

func getDp1(arr []int) []int {
	var dp = make([]int, len(arr), len(arr))
	for i := 0; i < len(arr); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] {
				dp[i] = int(math.Max(float64(dp[i]), float64(dp[j]+1)))
			}
		}
	}
	return dp
}

func GetLongestSubArr2(arr []int) []int {
	var res []int
	if len(arr) == 0 {
		return res
	}
	dp := getDp2(arr)
	return genLIS(arr, dp)
}

// getDP2 利用二分查找, 将时间复杂度控制在 O(NlogN)
func getDp2(arr []int) []int {
	var dp = make([]int, len(arr), len(arr))
	var ends = make([]int, len(arr), len(arr))
	var right int
	dp[0] = 1
	ends[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		l := 0
		r := right
		// 二分查找, 找到 最左边大于等于 arr[i] 的值
		for l <= r {
			m := (r + l) / 2
			if arr[i] > ends[m] {
				l = m + 1
			} else {
				r = m - 1
			}
		}
		right = int(math.Max(float64(l), float64(r)))
		ends[l] = arr[i]
		dp[i] = l + 1
	}
	return dp
}

func genLIS(arr, dp []int) []int {
	var l = 0
	var index = 0
	for i := 0; i < len(dp); i++ {
		if dp[i] > l {
			l = dp[i]
			index = i
		}
	}
	var lis = make([]int, l)
	lis[l-1] = arr[index]
	l -= 1
	for i := index; i >= 0; i-- {
		if arr[i] < arr[index] && dp[i] == dp[index]-1 {
			l -= 1
			lis[l] = arr[i]
			index = i
		}
	}
	return lis
}

// 动态规划 O(N^2)的实现2
// golang 这种切片支持的动态二维切片
// dp[i][] 表示以 arr[i] 结尾的最长递增子序列
// 计算出 dp , 获取最长的递增子序列即可

func LongestOrderArr(arr []int) []int {
	var l = len(arr)
	if l <= 1 {
		return arr
	}
	var dp = make([][]int, l, l)

	dp[0] = []int{arr[0]}
	for i := 1; i < l; i++ {
		tmp := []int{arr[i]}
		max := []int{}
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] && len(dp[j]) > len(max) {
				max = dp[j]
			}
		}
		dp[i] = append(max, tmp...)
	}
	// get long
	res := []int{}
	for i := 0; i < l; i++ {
		if len(dp[i]) > len(res) {
			res = dp[i]
		}
	}
	return res
}