package node_list

import (
	"fmt"

	"gitee.com/Anderson/my-algorithm/data_structure"
)

//打印两个有序链表的公共部分

func PrintCommon(head1, head2 *data_structure.Node) {
	var l1 = head1
	var l2 = head2
	for l1 != nil && l2 != nil {
		if l1.Value == l2.Value {
			fmt.Printf("common node. value:%v\n", l1.Value)
			l1 = l1.Next
			l2 = l2.Next
		} else if l1.Value < l2.Value {
			l1 = l1.Next
		} else {
			l2 = l2.Next
		}
	}

	return
}
