package node_list

import (
	"testing"

	"gitee.com/Anderson/my-algorithm/data_structure"
)

func getTestPlainNode() *data_structure.Node {
	var l1 = data_structure.NewNode(1)
	var l2 = data_structure.NewNode(2)
	var l3 = data_structure.NewNode(3)
	var l4 = data_structure.NewNode(2)
	var l5 = data_structure.NewNode(1)
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	return l1
}

func TestIsPalindrome(t *testing.T) {
	type args struct {
		head *data_structure.Node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{getTestPlainNode()},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.args.head); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
