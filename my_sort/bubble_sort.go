package my_sort

func BubbleSort(arr []int) []int {
	var l = len(arr)
	if l == 0 {
		return arr
	}
	for i := 0; i < l; i++ {
		for j := 0; j < l-i-1; j++ {
			if arr[j] < arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
	return arr
}


func BubbleSort1(arr []int) []int {
	var l = len(arr)
	if l == 0 {
		return arr
	}
	for i := 0; i <= l-1; i++ {
		for j := 0; j < l-i-1; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
	return arr
}
