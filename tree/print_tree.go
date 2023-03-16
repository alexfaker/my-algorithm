package tree

import (
	"math"

	"gitee.com/Anderson/my-algorithm/data_structure"
)

// 直观的打印二叉树

func PrintTree(head *data_structure.BinaryTreeNode) {

}

func getWidth(head *data_structure.BinaryTreeNode) int {
	if head == nil {
		return 0
	}
	return getWidth(head.Left) + 1 + getWidth(head.Right)
}

func getWeight(head *data_structure.BinaryTreeNode) int {
	if head == nil {
		return 0
	}
	maxWeight := 0
	if head.Left == nil && head.Right == nil {
		return 1
	} else {
		// 递归到叶子节点层，计算其宽度，然后迭代到根节点
		maxWeight = int(math.Max(float64(getWeight(head.Left)+getWeight(head.Right)), float64(maxWeight)))
	}
	return maxWeight
}

func getWeight2(head *data_structure.BinaryTreeNode) int {
	// todo 双端队列实现
	return 0
}


func getLongestLeft(head *data_structure.BinaryTreeNode) int {
	var res int
	return res
}
