package my_sort

// CountingSort
// @Description:  计数排序
// 条件限制, 假定给定的数组元素在 0-k之间
// 统计每个数字出现的次数
// @param arr
// @return []int
func CountingSort(arr []int, k int) []int {
	if len(arr) == 0 {
		return arr
	}
	var res = make([]int, k, k)
	for i := 0; i < len(arr); i++ {
		res[arr[i]] += 1
	}
	var ans = make([]int, 0, len(arr))
	for i, v := range res {
		if v == 0 {
			continue
		}
		for v > 0 {
			ans = append(ans, i)
			v--
		}
	}
	return arr
}
