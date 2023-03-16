package data_structure

type BinaryTreeNode struct {
	Value int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func NewBinaryTreeNode(val int) *BinaryTreeNode {
	return &BinaryTreeNode{Value: val}
}
