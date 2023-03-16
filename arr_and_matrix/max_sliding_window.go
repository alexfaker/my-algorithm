package arr_and_matrix

import "gitee.com/Anderson/my-algorithm/data_structure"

// 滑动窗口最大值
// 给你一个整数数组 nums，
// 有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。
// 你只可以看到在滑动窗口内的 k 个数字。
// 滑动窗口每次只向右移动一位。
// 返回 滑动窗口中的最大值 。

func GetMaxSlidingWindow(arr []int, w int) []int {
	var ans []int
	var l = len(arr)
	var left = 0
	var dqueue = data_structure.NewDoubleDirectQueueImp()
	for right := 0; right <= l-1; right++ {
		if right-left+1 == w { // 滑动窗口第一次完全展开
			ModifyDQueue(dqueue, left, right, arr)
			ans = append(ans, arr[dqueue.GetLeft()])

		} else if right-left == w { // 刚好超界, 收缩 1
			left++
			ModifyDQueue(dqueue, left, right, arr)
			ans = append(ans, arr[dqueue.GetLeft()])
		} else { // 还未完全展开
			ModifyDQueue(dqueue, left, right, arr)
		}

	}
	return ans
}

func ModifyDQueue(dq *data_structure.DoubleDirectQueueImp, l, r int, arr []int) {
	if dq.IsEmpty() {
		dq.RightIn(r)
		return
	}
	for !dq.IsEmpty() && arr[dq.GetRight()] < arr[r] {
		dq.RightOut()
	}
	dq.RightIn(r)
	// 检查左边
	if dq.GetLeft() < l {
		dq.LeftOut()
	}
	return
}
