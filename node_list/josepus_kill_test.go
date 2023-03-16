package node_list

import (
	"testing"

	"gitee.com/Anderson/my-algorithm/data_structure"
)

func getTestRingNode() *data_structure.Node {
	var l1 = data_structure.NewNode(1)
	var l2 = data_structure.NewNode(2)
	var l3 = data_structure.NewNode(3)
	var l4 = data_structure.NewNode(4)
	var l5 = data_structure.NewNode(5)
	var l6 = data_structure.NewNode(6)
	var l7 = data_structure.NewNode(7)
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	l5.Next = l6
	l6.Next = l7
	l7.Next = l1
	return l1
}

func TestJosePusKillV2(t *testing.T) {
	type args struct {
		head *data_structure.Node
		m    int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"tes1",
			args{
				head: getTestRingNode(),
				m:    1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := JosePusKillV2(tt.args.head, tt.args.m)
			println(got.Value)

		})
	}
}
