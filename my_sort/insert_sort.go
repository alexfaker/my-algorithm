package my_sort

func InsertSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	// 第 0 个位置默认已排好序, 从1 个位置开始遍历
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1 // j的 左边表示已经有序的部分
		// 找到当前在处理的值 arr[i] 应该插入的部分, 插入部分的右边元素依次往右移
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		// 在正确的位置插入, 保证 j 左边的有序
		arr[j+1] = key
	}
	return arr
}

func InsertSort1(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	// 第 0 个位置默认已排好序, 从1 个位置开始遍历
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1 // j的 左边表示已经有序的部分
		for j >= 0 && arr[j] < key {
			arr[j+1] = arr[j]
			// 找到当前在处理的值 arr[i] 应该插入的部分, 插入部分的右边元素依次往右移
			j = j - 1
		}
		// 在正确的位置插入, 保证 j 左边的有序
		arr[j+1] = key
	}
	return arr
}
