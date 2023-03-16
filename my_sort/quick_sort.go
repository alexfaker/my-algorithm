package my_sort

// QuickSort 快速排序
// @Description:  选定基数, 使得基数左边的位置的值小于于该基数, 右边位置的值大于该基数
// 然后递归的执行该操作
// @param arr
// @return []int
func QuickSort(arr []int) []int {
	var l = len(arr)
	if l == 0 {
		return arr
	}
	return quickSort(arr, 0, l-1)
}

func quickSort(arr []int, l, r int) []int {
	var i, j, t, tmp int
	if l > r {
		return arr
	}
	tmp = arr[l] // 基数
	i = l
	j = r
	for i != j {
		for arr[j] >= tmp && i < j {
			j--
		}
		for arr[i] <= tmp && i < j {
			i++
		}
		if i < j {
			t = arr[i]
			arr[i] = arr[j]
			arr[j] = t
		}
	}

	arr[l] = arr[i]
	arr[i] = tmp
	quickSort(arr, l, i-1)
	quickSort(arr, i+1, r)
	return arr
}
