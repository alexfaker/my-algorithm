package node_list

import (
	"testing"

	"gitee.com/Anderson/my-algorithm/data_structure"
)

func TestAddList1(t *testing.T) {
	type args struct {
		list1 *data_structure.Node
		list2 *data_structure.Node
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"test",
			args{
				list1: getTestNode(),
				list2: getTestPlainNode(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			 got := AddList1(tt.args.list1, tt.args.list2)
			 plnNodeList(got)
		})
	}
}
