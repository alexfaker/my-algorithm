package tree

import (
	"fmt"
	"math"

	"gitee.com/Anderson/my-algorithm/data_structure"
)

// 打印二叉树的边界
// 头节点和叶子结点为边界节点
// 如果节点在其所在的层中是最左或者最右的, 也是边界结点
func PrintEdge(head *data_structure.BinaryTreeNode) {
	if head == nil {
		return
	}
	// 获取树的高度
	l := getHeight(head, 0)
	// 找到树每一层的左右边界
	var edgeMap = make([][2]*data_structure.BinaryTreeNode, l, l)
	getEdgeMap(head, 0, edgeMap)
	// 打印左边界
	for i := 0; i < l; i++ {
		fmt.Printf(" %d ", edgeMap[i][0].Value)
	}
	// 打印非边界的叶子结点
	printLeafNotEdge(head, 0, edgeMap)
	// 打印右边界
	for i := l - 1; i >= 0; i-- {
		if edgeMap[i][0].Value != edgeMap[i][1].Value {
			fmt.Printf(" %d ", edgeMap[i][1].Value)
		}
	}
}

func getHeight(head *data_structure.BinaryTreeNode, l int) int {
	if head == nil {
		return l
	}
	return int(math.Max(float64(getHeight(head.Left, l+1)), float64(getHeight(head.Right, l+1))))
}

func getEdgeMap(head *data_structure.BinaryTreeNode, l int, edgeMap [][2]*data_structure.BinaryTreeNode) {
	if head == nil {
		return
	}
	if edgeMap[l][0] == nil {
		edgeMap[l][0] = head
	}
	edgeMap[l][1] = head
	getEdgeMap(head.Left, l+1, edgeMap)
	getEdgeMap(head.Right, l+1, edgeMap)
}

func printLeafNotEdge(head *data_structure.BinaryTreeNode, l int, edgeMap [][2]*data_structure.BinaryTreeNode) {
	if head == nil {
		return
	}
	if head.Left == nil && head.Right == nil &&
		edgeMap[l][0].Value != head.Value &&
		edgeMap[l][1].Value != head.Value {
		fmt.Printf("%d ", head.Value)
	}
	printLeafNotEdge(head.Left, l+1, edgeMap)
	printLeafNotEdge(head.Right, l+1, edgeMap)
}
