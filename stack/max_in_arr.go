package stack

// 生成窗口最大值数组
//如何写出O(N)的算法？
//思路：用一个辅助数据结构（双端队列 || 数组）记录当前窗口的最大值， 在遍历过程中 记录每个窗口的最大值
//func main() {
//	var arr = []int{4, 3, 5, 4, 3, 3, 6, 7}
//	res := getMaxWindow(arr, 3)
//	fmt.Println(res)
//}

func getMaxWindow(arr []int, w int) []int {
	var res = make([]int, len(arr)-w+1, len(arr)-w+1)
	if len(arr) <= 0 || w < 1 {
		return res
	}
	var qmax = make([]int, 0, 0)

	var index = 0
	for i := 0; i < len(arr); i++ {
		for len(qmax) > 0 && arr[qmax[len(qmax)-1]] <= arr[i] {
			qmax = qmax[:len(qmax)-1]
		}
		qmax = append(qmax, i)
		if qmax[0] == i-w {
			qmax = qmax[1 : len(qmax)-1]
		}
		if i >= w-1 {
			res[index] = arr[qmax[0]]
			index++
		}
	}
	return res
}
