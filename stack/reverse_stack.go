package stack

import (
	"gitee.com/Anderson/my-algorithm/data_structure"
)

// 只用递归函数 逆序一个栈
func ReverseStack1(myStack *data_structure.Stack) {
	if myStack.IsEmpty() {
		return
	}
	res := getAndRmLastElement(myStack)
	ReverseStack1(myStack)
	myStack.Push(res)
	return
}

func getAndRmLastElement(myStack *data_structure.Stack) int {
	var res, _ = myStack.Pop()
	if myStack.IsEmpty() {
		return res
	} else {
		last := getAndRmLastElement(myStack)
		myStack.Push(res)
		return last
	}
}

func ReverseStack(myStack *data_structure.Stack) {
	if myStack.IsEmpty() {
		return
	}
	var i = getAndRmLastElement(myStack)
	ReverseStack(myStack)
	myStack.Push(i)
}
