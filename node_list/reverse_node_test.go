package node_list

import (
	"fmt"
	"reflect"
	"testing"

	"gitee.com/Anderson/my-algorithm/data_structure"
)

func getTestNode() *data_structure.Node {
	var l1 = data_structure.NewNode(1)
	var l2 = data_structure.NewNode(2)
	var l3 = data_structure.NewNode(3)
	var l4 = data_structure.NewNode(5)
	var l5 = data_structure.NewNode(7)
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	return l1
}
func TestReverseNode(t *testing.T) {
	type args struct {
		head *data_structure.Node
	}
	var l1 = data_structure.NewNode(1)
	var l2 = data_structure.NewNode(2)
	var l3 = data_structure.NewNode(3)
	var l4 = data_structure.NewNode(5)
	var l5 = data_structure.NewNode(7)
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	tests := []struct {
		name string
		args args
		want *data_structure.Node
	}{
		{
			"test1",
			args{head: l1},
			l5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseNode(tt.args.head)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseNode() = %v, want %v", got, tt.want)
			}
			t.Logf("new head:%v", got.Value)
		})
	}
}

func TestReversePartNode(t *testing.T) {
	type args struct {
		head *data_structure.Node
		from int
		to   int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"test1",
			args{getTestNode(), 2, 5},
		},
		{
			"test1",
			args{getTestNode(), 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReversePartNode(tt.args.head, tt.args.from, tt.args.to)
			plnNodeList(got)
		})
	}
}

func plnNodeList(head *data_structure.Node) {
	fmt.Println("pln nodehead:")
	for head != nil {
		fmt.Println(head.Value)
		head = head.Next
	}

}
