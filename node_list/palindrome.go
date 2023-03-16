package node_list

import (
	"gitee.com/Anderson/my-algorithm/data_structure"
)

//判断单链表是否属于回文结构
//比如
//1->2->3->2->1 是回文结构
//1->2->2->1 是回文结构
//1->2->3 不是回文结构
//要求 时间复杂度O(N), 空间复杂度O(1)

//简单的实现， 利用数据结构 栈， 将链表压入栈 再出栈（相当于逆序了），
//如果出栈时各个节点的值与原链表一一对应， 说明是回文结构
//时间复杂度O(N), 空间复杂度O(N)

func IsPalindrome(head *data_structure.Node) bool {
	if head.Next == nil {
		return true
	}
	var myStack = data_structure.NewStack()
	for cur := head; cur != nil; cur = cur.Next {
		myStack.Push(cur.Value)
	}

	//出栈， 判断值是否一一对应
	var isPal = true
	for cur := head; cur != nil; cur = cur.Next {
		res, err := myStack.Pop()
		if err != nil {
			isPal = false
			break
		}
		if res != cur.Value {
			isPal = false
			break
		}
	}
	return isPal
}

//时间复杂度O(1)的实现

func IsPalindrome2(head *data_structure.Node) bool {
	if head.Next == nil {
		return true
	}
	var isPalindrome = true
	var flow = head
	var fast = head.Next
	for flow != nil && fast != nil {
		flow = flow.Next
		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			fast = nil
		}
	}
	var middle = flow
	//逆序右半部分的链表
	rHead := ReverseNode(flow)

	var l1 = head
	var l2 = rHead
	for l1 != nil && l2 != nil {
		if l1.Value != l2.Value {
			isPalindrome = false
			break
		} else {
			l1 = l1.Next
			l2 = l2.Next
		}
	}

	//恢复链表
	rrHead := ReverseNode(rHead)
	middle.Next = rrHead.Next
	return isPalindrome
}
