package tree

import (
	"testing"

	"gitee.com/Anderson/my-algorithm/data_structure"
)

func getTestEdgeTree() *data_structure.BinaryTreeNode {
	b1 := data_structure.NewBinaryTreeNode(1)
	b2 := data_structure.NewBinaryTreeNode(2)
	b3 := data_structure.NewBinaryTreeNode(3)
	b4 := data_structure.NewBinaryTreeNode(4)
	b5 := data_structure.NewBinaryTreeNode(5)
	b6 := data_structure.NewBinaryTreeNode(6)
	b7 := data_structure.NewBinaryTreeNode(7)
	b8 := data_structure.NewBinaryTreeNode(8)
	b9 := data_structure.NewBinaryTreeNode(9)
	b10 := data_structure.NewBinaryTreeNode(10)
	b11 := data_structure.NewBinaryTreeNode(11)
	b12 := data_structure.NewBinaryTreeNode(12)
	b13 := data_structure.NewBinaryTreeNode(13)
	b14 := data_structure.NewBinaryTreeNode(14)
	b15 := data_structure.NewBinaryTreeNode(15)
	b16 := data_structure.NewBinaryTreeNode(16)
	b1.Left = b2
	b1.Right = b3
	b2.Right = b4
	b3.Left = b5
	b3.Right = b6
	b4.Left = b7
	b4.Right = b8
	b5.Left = b9
	b5.Right = b10
	b8.Right = b11
	b9.Left = b12
	b11.Left = b13
	b11.Right = b14
	b12.Left = b15
	b12.Right = b16
	return b1
}

func TestPrintEdge(t *testing.T) {
	type args struct {
		head *data_structure.BinaryTreeNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"1",
			args{getTestEdgeTree()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintEdge(tt.args.head)
		})
	}
}
