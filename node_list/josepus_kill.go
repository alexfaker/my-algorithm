package node_list

import (
	"gitee.com/Anderson/my-algorithm/data_structure"
)

//单向环形链表模拟约瑟夫问题
//所有人拍成一个圆圈， 从第一个人开始报数 1 ， 第二个人 2 ， 第三个人 3， 报到 3 的人立刻自杀，
//下一个人重新从 1 开始报数 ， 报到 3 的人自杀， 最后只剩下一个人

// 普通解法 时间复杂度O(N*m)

func JosePusKill(head *data_structure.Node, m int) *data_structure.Node {
	var live = head
	var pre = head
	var curNum = 1
	if m == 1 || head.Next == nil {
		return nil
	}
	for live.Next != live {
		if curNum == m {
			pre.Next = live.Next
			curNum = 1
			live = live.Next
		} else {
			curNum++
			pre = live
			live = live.Next
		}
	}
	return live
}

func JosePusKillV2(head *data_structure.Node, m int) *data_structure.Node {
	var live = head
	var pre *data_structure.Node
	var curNum = 1
	if head.Next == nil || m == 1 {
		live = head
		return live
	}

	for live.Next != live {
		if curNum == m {
			curNum = 1
			pre.Next = live.Next
			live = live.Next
		} else {
			pre = live
			live = live.Next
			curNum++
		}
	}
	return live
}

// 进阶解法 如何实现时间复杂度O(N) 的算法？ todo

func JosePusKillImprove(head *data_structure.Node, m int) *data_structure.Node {
	var live = head
	var pre = head
	var curNum = 1
	if m == 1 || head.Next == nil {
		return nil
	}
	for live.Next != live {
		if curNum == m {
			pre.Next = live.Next
			curNum = 1
			live = live.Next
		} else {
			curNum++
			pre = live
			live = live.Next
		}
	}
	return live
}
