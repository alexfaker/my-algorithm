package node_list

import "gitee.com/Anderson/my-algorithm/data_structure"

// 两个单链表生成相加链表

func AddList1(list1, list2 *data_structure.Node) *data_structure.Node {
	var stack1 = data_structure.NewStack()
	var stack2 = data_structure.NewStack()
	tmp1 := list1
	for tmp1 != nil {
		stack1.Push(tmp1.Value)
		tmp1 = tmp1.Next
	}

	tmp2 := list2
	for tmp2 != nil {
		stack2.Push(tmp2.Value)
		tmp2 = tmp2.Next
	}
	// gen res list
	var overflow int
	var res *data_structure.Node
	for !stack1.IsEmpty() || !stack2.IsEmpty() {
		t1 := 0
		t2 := 0
		if !stack1.IsEmpty() {
			t1, _ = stack1.Pop()
		}
		if !stack2.IsEmpty() {
			t2, _ = stack2.Pop()
		}
		tmp := (t1 + t2 + overflow) % 10
		overflow = (t1 + t2 + overflow) / 10
		newNode := data_structure.NewNode(tmp)
		if res != nil {
			newNode.Next = res
		}
		res = newNode
	}
	return res
}
