package stack

import (
	"gitee.com/Anderson/my-algorithm/data_structure"
)

type MinStack interface {
	Push(i int)
	Pop() (int, error)
	GetMin() (int, error)
}

type MinStackImp struct {
	NormalStack data_structure.Stack
	MinStack    data_structure.Stack
}

func NewMinStackImp() *MinStackImp {
	return &MinStackImp{NormalStack: data_structure.Stack{List: []int{}},
		MinStack: data_structure.Stack{List: []int{}}}
}

func (m *MinStackImp) Push(i int) {
	m.NormalStack.Push(i)
	if m.MinStack.IsEmpty() {
		m.MinStack.Push(i)
	} else {
		// 比较当前元素与minStack栈顶， 如果该值小于等于栈顶元素， 入栈
		res, _ := m.MinStack.Peek()
		if i <= res {
			m.MinStack.Push(i)
		}
	}
	return
}

func (m *MinStackImp) Pop() (int, error) {
	var res int
	var err error
	res, err = m.NormalStack.Pop()
	if err != nil {
		return res, err
	}
	minRes, err := m.MinStack.Peek()
	if err != nil {
		return res, err
	}
	if res <= minRes {
		m.MinStack.Pop()
	}
	return res, err
}

func (m *MinStackImp) GetMin() (int, error) {
	return m.MinStack.Peek()
}
