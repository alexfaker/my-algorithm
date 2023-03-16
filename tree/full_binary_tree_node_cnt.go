package tree

import "gitee.com/Anderson/my-algorithm/data_structure"

/*
统计完全二叉树的节点数
给定一颗完全二叉树的头节点 head, 返回这棵树的节点个数
如果节点数为 N, 请实现时间复杂度低于 O(N)的解法

完全二叉树的含义(from 百度百科)
一棵深度为k的有n个结点的二叉树，对树中的结点按从上至下、从左到右的顺序进行编号，
如果编号为i（1≤i≤n）的结点与满二叉树中编号为i的结点在二叉树中的位置相同，
则这棵二叉树称为完全二叉树
*/

/*
递归实现的代码就是这么简洁, 但是实际要搞懂还是需要自己多多练习一下子
*/

func CBTNodeNum(head *data_structure.BinaryTreeNode) int {
	if head == nil {
		return 0
	}
	return bs(head, 1, mostLeftLevel(head, 1))
}

func bs(node *data_structure.BinaryTreeNode, curLevel, height int) int {
	if curLevel == height {
		return 1
	}
	if mostLeftLevel(node.Right, curLevel+1) == height {
		// 如果 节点的右子树的高度 ==树高, 说明节点的左子树一定是满二叉树, 满二叉树的节点个数为 2^h -1
		// 加上该节点自己 就是 2^h - 1 + 1
		// 整棵树的节点数 就是 左子树节点个数 + 节点自己 + 节点右子树个数(递归)
		return 1<<(height-curLevel) + bs(node.Right, curLevel+1, height)
	} else {
		// 如果 节点的右子树的高度 !=树高, 说明节点的右子树一定是满二叉树, 左子树满不满不确定
		// 整棵树的节点数 就是 右子树节点个数 + 节点自己 + 节点左子树个数(递归)
		return 1<<(height-1-curLevel) + bs(node.Left, curLevel+1, height)
	}

}

func mostLeftLevel(node *data_structure.BinaryTreeNode, level int) int {
	for node != nil {
		level += 1
		node = node.Left
	}
	return level - 1
}
