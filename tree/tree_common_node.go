package tree

import (
	"gitee.com/Anderson/my-algorithm/data_structure"
)

/*
给定一颗二叉树,头节点为 head
节点 o1, o2 .这两个节点一定在二叉树上
寻找两个节点的最低公共祖先

最低公共祖先:两个节点都往上找, 第一个相交的节点即为最低公共祖先
*/

/*
动态规划解法
向左子树 右子树 要o1, o2  这两个节点
1. 如果左子树 不为空 && 右子树不为空, 那么 当前节点就是 o1, o2 的最小公共祖先
2. 如果左子树不为空, 右子树为空, 那么 o1, o2 都在左子树上, 当前左子树返回的节点就是 最小公共祖先
3. 如果右子树不为空, 左子树为空, 那么 o1, o2 都在右子树上, 当前右子树返回的节点就是 最小公共祖先
4. 如果都为空, 那就是入参有问题, 不符合题意(o1, o2 都在二叉树上)
*/

func LowestAncestor(head, o1, o2 *data_structure.BinaryTreeNode) *data_structure.BinaryTreeNode {
	if head == nil || head == o1 || head == o2 {
		return head
	}
	left := LowestAncestor(head.Left, o1, o2)
	right := LowestAncestor(head.Left, o1, o2)
	if left != nil && right != nil {
		return head
	} else if left != nil {
		return left
	} else {
		return right
	}
}

/*
 额外空间的解法 , 使用一种数据结构记录每一个节点的父节点
*/

func LowestAncestor1(head, o1, o2 *data_structure.BinaryTreeNode) *data_structure.BinaryTreeNode {
	//特殊情况, O1 和 o2 直接相邻
	if (o1.Left != nil && o1.Left == o2) || (o1.Right != nil && o1.Right == o2) {
		return o1
	}
	if (o2.Left != nil && o2.Left == o1) || (o2.Right != nil && o2.Right == o1) {
		return o2
	}

	// 记录每个节点的父节点
	var pMap = make(map[*data_structure.BinaryTreeNode]*data_structure.BinaryTreeNode)
	pMap[head] = nil
	process(head, pMap)

	var o1Map = make(map[*data_structure.BinaryTreeNode]bool)
	o1Map[o1] = true
	cur, ok := pMap[o1]
	if !ok {
		return nil
	}
	// o1 上浮的 路径
	for cur != nil {
		o1Map[cur] = true
		cur = pMap[cur]
	}

	cur1 := o2
	for cur1 != nil {
		if _, ok := o1Map[cur1]; ok {
			return cur1
		} else {
			cur1 = pMap[cur1]
		}
	}
	return nil
}

func process(head *data_structure.BinaryTreeNode, pMap map[*data_structure.BinaryTreeNode]*data_structure.BinaryTreeNode) {
	if head == nil {
		return
	}
	if head.Left != nil {
		pMap[head.Left] = head
		process(head.Left, pMap)
	}
	if head.Right != nil {
		pMap[head.Right] = head
		process(head.Right, pMap)
	}

}
