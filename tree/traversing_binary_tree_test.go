package tree

import (
	"testing"

	"gitee.com/Anderson/my-algorithm/data_structure"
)

func getTestTree() *data_structure.BinaryTreeNode {
	var b1 = data_structure.NewBinaryTreeNode(3)
	var b2 = data_structure.NewBinaryTreeNode(2)
	var b3 = data_structure.NewBinaryTreeNode(1)
	var b4 = data_structure.NewBinaryTreeNode(5)
	var b5 = data_structure.NewBinaryTreeNode(6)
	var b6 = data_structure.NewBinaryTreeNode(4)
	var b7 = data_structure.NewBinaryTreeNode(9)
	b1.Left = b2
	b1.Right = b3
	b2.Left = b4
	b2.Right = b5
	b3.Left = b6
	b3.Right = b7
	return b1
}

func TestPreOrder(t *testing.T) {
	type args struct {
		head *data_structure.BinaryTreeNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"1",
			args{getTestTree()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PreOrder(tt.args.head)
		})
	}
}

func TestMidOrder(t *testing.T) {
	type args struct {
		head *data_structure.BinaryTreeNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"1",
			args{getTestTree()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MidOrder(tt.args.head)
		})
	}
}

func TestAfterOrder(t *testing.T) {
	type args struct {
		head *data_structure.BinaryTreeNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"1",
			args{getTestTree()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AfterOrder(tt.args.head)
		})
	}
}

func TestMorrisPre(t *testing.T) {
	type args struct {
		head *data_structure.BinaryTreeNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"1",
			args{getTestSerialTree()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MorrisPre(tt.args.head)
		})
	}
}

func TestMorrisIn(t *testing.T) {
	type args struct {
		head *data_structure.BinaryTreeNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"1",
			args{getTestSerialTree()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MorrisIn(tt.args.head)
		})
	}
}
