package arr_and_matrix

import (
	"math"
)

/*
 给定一个无序数组 arr, 每个元素均为正值, 在给定一个正数 k
求 arr 的所有子数组中元素相加和为 k 的最长子数组长度
例如 arr= [1,2,1,1,1] k= 3
累加和为 3 的最长子数组为 [1,1,1], 结果返回3
解法: 滑动窗口
*/

func GetMaxLength(arr []int, k int) int {
	var l = len(arr)
	if l == 0 || k <= 0 {
		return 0
	}
	var left int
	var right int
	var maxLen int
	var sum = arr[0]
	for right < l {
		if sum == k {
			maxLen = int(math.Max(float64(maxLen), float64(right-left+1)))
			sum -= arr[left]
			left++
		} else if sum < k {
			right++
			if right == l {
				break
			}
			sum += arr[right]
		} else {
			sum -= arr[left]
			left++
		}

	}
	return maxLen

}

/*
扩展, 数组元素可以为任意 负数, 正数, 0
比较 low 的写法, 时间复杂度 O(N*N) ,空间复杂度O(1)
*/
func GetMaxLength1(arr []int, k int) int {
	var l = len(arr)
	if l == 0 || k <= 0 {
		return 0
	}
	var maxLen int
	for left := 0; left < l; left++ {
		sum := arr[left]
		if sum == k {
			maxLen = int(math.Max(float64(maxLen), float64(1)))
		}
		for right := left + 1; right < l; right++ {
			sum += arr[right]
			if sum == k {
				maxLen = int(math.Max(float64(maxLen), float64(right-left+1)))
			}
		}
	}
	return maxLen

}

/*
扩展, 数组元素可以为任意 负数, 正数, 0
比较 high 的写法, 时间复杂度 O(N) ,空间复杂度O(N)
这个我看是看懂了, 但是为什么这样做, 还是有一点疑问???
  1. 为什么 map 只记录 sum 的第一个节点呢?
*/

func GetMaxLength2(arr []int, k int) int {
	var l = len(arr)
	if l == 0 || k <= 0 {
		return 0
	}
	var maxLen int
	var sumMap = map[int]int{0: -1}
	var sum = 0
	for i := 0; i < l; i++ {
		sum += arr[i]
		if _, ok := sumMap[sum-k]; ok {
			maxLen = int(math.Max(float64(maxLen), float64(i-sumMap[sum-k])))
		}
		if _, ok := sumMap[sum]; !ok {
			sumMap[sum] = i
		}
	}
	return maxLen

}
