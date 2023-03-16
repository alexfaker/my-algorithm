package my_sort

/*
归并排序
时间复杂度 O(NlogN)
空间复杂度O(NlogN)
*/

func MergeSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	return mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int, left, right int) []int {
	if left == right {
		return []int{arr[left]}
	}
	// 左边排序,
	// 右边排序
	// 归并两个排序的结果
	var mid = (left + right) / 2
	leftArr := mergeSort(arr, left, mid)
	rightArr := mergeSort(arr, mid+1, right)
	return merge(leftArr, rightArr)
}

func merge(leftArr, rightArr []int) []int {
	var res []int
	var l1 = len(leftArr)
	var l2 = len(rightArr)
	if l1 == 0 {
		res = rightArr
		return res
	} else if l2 == 0 {
		res = leftArr
		return res
	}

	var i = 0
	var j = 0
	for i < l1 && j < l2 {
		if leftArr[i] <= rightArr[j] {
			res = append(res, leftArr[i])
			i++
		} else {
			res = append(res, rightArr[j])
			j++
		}
	}
	if i < l1 {
		res = append(res, leftArr[i:]...)
	}
	if j < l2 {
		res = append(res, rightArr[j:]...)
	}
	return res
}


func MergeSort1(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	return mergeSort1(arr, 0, len(arr)-1)
}

func mergeSort1(arr []int, left, right int) []int {
	if left == right {
		return []int{arr[left]}
	}
	// 左边排序,
	// 右边排序
	// 归并两个排序的结果
	var mid = (left + right) / 2
	leftArr := mergeSort1(arr, left, mid)
	rightArr := mergeSort1(arr, mid+1, right)
	return merge1(leftArr, rightArr)
}

func merge1(leftArr, rightArr []int) []int {
	var res []int
	var l1 = len(leftArr)
	var l2 = len(rightArr)
	if l1 == 0 {
		res = rightArr
		return res
	} else if l2 == 0 {
		res = leftArr
		return res
	}

	var i = 0
	var j = 0
	for i < l1 && j < l2 {
		if leftArr[i] > rightArr[j] {
			res = append(res, leftArr[i])
			i++
		} else {
			res = append(res, rightArr[j])
			j++
		}
	}
	if i < l1 {
		res = append(res, leftArr[i:]...)
	}
	if j < l2 {
		res = append(res, rightArr[j:]...)
	}
	return res
}
