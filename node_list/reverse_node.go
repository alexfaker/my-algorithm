package node_list

import (
	"gitee.com/Anderson/my-algorithm/data_structure"
)

//反转单向链表
func ReverseNode(head *data_structure.Node) *data_structure.Node {
	var pre *data_structure.Node
	var next = head

	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

// 翻转部分单向链表  from to
func ReversePartNode(head *data_structure.Node, from, to int) *data_structure.Node {
	var nl int
	var p = head
	var fpre, tpos *data_structure.Node
	for p != nil {
		nl++
		if from-1 == nl {
			fpre = p
		}
		if nl == to+1 {
			tpos = p
		}
		p = p.Next
	}
	if from < 1 || to > nl || from > to {
		return head
	}

	var node1 = fpre.Next
	var node2 = node1.Next
	node1.Next = tpos

	var next *data_structure.Node
	for node2 != tpos {
		next = node2.Next
		node2.Next = node1
		node1 = node2
		node2 = next
	}
	if fpre != nil {
		fpre.Next = node1
		return head
	}
	return node1
}
