package stack

import (
	"gitee.com/Anderson/my-algorithm/data_structure"
)

/*
复制含有随机指针节点的链表
一种特殊的单链表节点如下
rand 指针可指向链表任意位置， 或者为nil
给定一个由randNode节点组成的无环单链表头节点head，请完成一个复制函数， 返回新链表的头节点
要求 时间复杂度O(N)， 空间复杂度O(1)
*/

func CopyRandNode(head *data_structure.RandStackNode) *data_structure.RandStackNode {
	if head == nil {
		return nil
	}
	cur := head
	for cur != nil {
		next := cur.Next
		newNode := data_structure.NewRandStackNode(cur.Value)
		cur.Next = newNode
		newNode.Next = next

		cur = newNode.Next
	}
	//依据新老节点的相对位置， 确定新节点的随机指针指向
	cur = head
	for cur != nil {
		newNode := cur.Next
		next := cur.Next.Next
		newNode.Rand = cur.Rand.Next

		cur = next
	}
	//分离新老链表
	cur = head
	var newHead = cur.Next
	for cur != nil {
		newNode := cur.Next
		next := cur.Next.Next
		cur.Next = next
		if next != nil {
			newNode.Next = next.Next
		} else {
			newNode.Next = nil
		}

		cur = next
	}
	return newHead
}
