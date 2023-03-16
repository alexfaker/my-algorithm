package stack

import (
	"gitee.com/Anderson/my-algorithm/data_structure"
)


type MyQueue interface {
	Add(i int)
	Poll() (int, error)
	Peek() (int, error)
}

type QueueByStack struct {
	QStack1 data_structure.Stack
	QStack2 data_structure.Stack
}

func NewQueueByStack() *QueueByStack {
	return &QueueByStack{QStack1: data_structure.Stack{List: []int{}}, QStack2: data_structure.Stack{List: []int{}}}
}

func (q *QueueByStack) Add(i int) {
	q.QStack1.Push(i)
}

func (q *QueueByStack) Poll() (int, error) {
	var res int
	if !q.QStack2.IsEmpty() {
		return q.QStack2.Pop()
	}
	for !q.QStack1.IsEmpty() {
		tmp, err := q.QStack1.Pop()
		if err != nil {
			return res, err
		}
		q.QStack2.Push(tmp)
	}
	return q.QStack2.Pop()
}

func (q *QueueByStack) Peek() (int, error) {
	var res int
	if !q.QStack2.IsEmpty() {
		return q.QStack2.Pop()
	}
	for !q.QStack1.IsEmpty() {
		tmp, err := q.QStack1.Pop()
		if err != nil {
			return res, err
		}
		q.QStack2.Push(tmp)
	}
	return q.QStack2.Peek()
}

func (q *QueueByStack) IsEmpty() bool {
	return q.QStack1.IsEmpty() && q.QStack2.IsEmpty()
}
