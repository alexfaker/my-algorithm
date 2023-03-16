package tree

import (
	"fmt"
	"testing"

	"gitee.com/Anderson/my-algorithm/data_structure"
)

func TestPrintTree(t *testing.T) {
	type args struct {
		head *data_structure.BinaryTreeNode
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestGetWidth(t *testing.T) {
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
			fmt.Println(getWeight(tt.args.head))
		})
	}
}
