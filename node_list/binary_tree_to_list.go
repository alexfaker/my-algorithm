package node_list

import (
	"gitee.com/Anderson/my-algorithm/data_structure"
	"gitee.com/Anderson/my-algorithm/stack"
)

//搜索二叉树转换成双向链表

//ps 二叉树 的数据结构除了值外，还有左子树、右子树两个指针， 和双向链表结构上有相似性
//可以转换

//要求： 中序遍历二叉树，转换成双向链表

func Convert1(head *data_structure.BinaryTreeNode) *data_structure.BothWayNode {
	if head == nil {
		return nil
	}
	if head.Left == nil && head.Right == nil {
		return data_structure.NewBothWayNode(head.Value)
	}

	var queue = stack.NewQueueByStack()
	MiddleIterToList(head, queue)
	h, err := queue.Poll()
	if err != nil {
		panic(err)
	}
	var newHead = data_structure.NewBothWayNode(h)
	pre := newHead
	for !queue.IsEmpty() {
		res, _ := queue.Poll()
		cur := data_structure.NewBothWayNode(res)
		pre.Next = cur
		cur.Last = pre
		pre = cur
	}
	return newHead
}

func MiddleIterToList(head *data_structure.BinaryTreeNode, queue *stack.QueueByStack) {
	if head == nil {
		return
	}
	MiddleIterToList(head.Left, queue)
	queue.Add(head.Value)
	MiddleIterToList(head.Right, queue)
}
