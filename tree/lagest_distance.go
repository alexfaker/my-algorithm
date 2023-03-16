package tree

import (
	"math"

	"gitee.com/Anderson/my-algorithm/data_structure"
)

/*
从二叉树的节点A 出发, 可以向上或者向下走, 但沿途的节点只能经过一次
当到点节点 B 时, 路径上的节点数叫做 A 到 B 的距离

给定一个二叉树 , 头节点为 head, 求 整棵树上节点间的最大距离
*/

/*
可以用动态规划搞定, 对于头节点 h, 求出左子树到头节点的最大距离, 右子树也一样
那么整棵树的最长距离必然是其中一种
1.  h 的左子树上的最大距离(注意, 是左子树上的最大路径, 不一定包含h.left 这个节点)
2. h 的右子树上的最大距离(同理)
3. h左子树到h.left的最远距离 + 1(代表 h) + h右子树到h.right的最远距离 ==> 简称 l+1+r

每次递归的时候, 将这三者中的最大值返回即可
*/

func MaxDistance(node *data_structure.BinaryTreeNode) int {
	lr, subMax := calMaxDistance(node)
	return int(math.Max(float64(lr), float64(subMax)))
}
func calMaxDistance(node *data_structure.BinaryTreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}
	// 左子树上的最大距离,左子树上距离 node 最远的叶节点的记录
	leftMax, maxFromLeft := calMaxDistance(node.Left)
	rightMax, maxFromRight := calMaxDistance(node.Right)
	// l+1+r
	curNodeMax := maxFromLeft + maxFromRight + 1
	return int(math.Max(float64(curNodeMax), math.Max(float64(leftMax), float64(rightMax)))), int(math.Max(float64(maxFromLeft), float64(maxFromRight))) + 1
}
