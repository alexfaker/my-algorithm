package data_structure

import (
	"errors"
)

type RandStackNode struct {
	Value int
	Next  *RandStackNode
	Rand  *RandStackNode
}

func NewRandStackNode(val int) *RandStackNode {
	return &RandStackNode{Value: val}
}

type Stack struct {
	List []int
}

func NewStack() *Stack {
	return &Stack{List: []int{}}
}

func (s *Stack) Push(i int) {
	s.List = append(s.List, i)
}

func (s *Stack) Pop() (int, error) {
	var res int
	var err error
	if len(s.List) > 0 {
		res = s.List[len(s.List)-1]
		s.List = s.List[:len(s.List)-1]
	} else {
		return -1, errors.New("stack is empty")
	}
	return res, err
}

func (s *Stack) IsEmpty() bool {
	return len(s.List) == 0
}

func (s *Stack) Peek() (int, error) {
	var res int
	var err error
	if len(s.List) > 0 {
		res = s.List[len(s.List)-1]
	} else {
		return -1, errors.New("stack is empty")
	}
	return res, err
}

func (s *Stack) Size() int {
	return len(s.List)
}

type AllStack struct {
	List []interface{}
}

func NewAllStack() *AllStack {
	return &AllStack{List: []interface{}{}}
}

func (s *AllStack) Push(i interface{}) {
	s.List = append(s.List, i)
}

func (s *AllStack) Pop() (interface{}, error) {
	var res interface{}
	var err error
	if len(s.List) > 0 {
		res = s.List[len(s.List)-1]
		s.List = s.List[:len(s.List)-1]
	} else {
		return -1, errors.New("stack is empty")
	}
	return res, err
}

func (s *AllStack) IsEmpty() bool {
	return len(s.List) == 0
}

func (s *AllStack) Peek() (interface{}, error) {
	var res interface{}
	var err error
	if len(s.List) > 0 {
		res = s.List[len(s.List)-1]
	} else {
		return -1, errors.New("stack is empty")
	}
	return res, err
}

func (s *AllStack) Size() int {
	return len(s.List)
}

type BothWayNode struct {
	Value int
	Next  *BothWayNode
	Last  *BothWayNode
}

func NewBothWayNode(value int) *BothWayNode {
	return &BothWayNode{Value: value}
}

type Node struct {
	Value int
	Next  *Node
}

func NewNode(value int) *Node {
	return &Node{Value: value}
}
