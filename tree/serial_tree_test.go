package tree

import (
	"fmt"
	"testing"

	"gitee.com/Anderson/my-algorithm/data_structure"
)

func getTestSerialTree() *data_structure.BinaryTreeNode {
	var b1 = data_structure.NewBinaryTreeNode(1)
	var b2 = data_structure.NewBinaryTreeNode(2)
	var b3 = data_structure.NewBinaryTreeNode(3)
	var b4 = data_structure.NewBinaryTreeNode(4)
	var b5 = data_structure.NewBinaryTreeNode(5)
	var b6 = data_structure.NewBinaryTreeNode(6)
	b1.Left = b2
	b1.Right = b3
	b2.Left = b4
	b2.Right = b5
	b3.Left = b6
	return b1
}

func TestPreSerial(t *testing.T) {
	type args struct {
		head *data_structure.BinaryTreeNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{getTestSerialTree()},
			"1!2!4!#!#!5!#!#!3!6!#!#!#!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PreSerial(tt.args.head); got != tt.want {
				t.Errorf("PreSerial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRevertPreSerial(t *testing.T) {
	type args struct {
		res string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"1",
			args{"1!2!4!#!#!5!#!#!3!6!#!#!#!"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RevertPreSerial(tt.args.res)
			fmt.Println(PreSerial(got))
		})
	}
}

func TestLevelSerial(t *testing.T) {
	type args struct {
		head *data_structure.BinaryTreeNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{getTestSerialTree()},
			"1!2!3!4!5!6!#!#!#!#!#!#!#!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LevelSerial(tt.args.head); got != tt.want {
				t.Errorf("LevelSerial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRevertLevelSerial(t *testing.T) {
	type args struct {
		res string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"1",
			args{"1!2!3!4!5!6!#!#!#!#!#!#!#!"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RevertLevelSerial(tt.args.res)
			fmt.Println(LevelSerial(got))

		})
	}
}
