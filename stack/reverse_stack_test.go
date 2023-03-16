package stack

import (
	"fmt"
	"testing"

	"gitee.com/Anderson/my-algorithm/data_structure"
)

func GetTestStack() *data_structure.Stack {
	var s1 = data_structure.NewStack()
	s1.Push(1)
	s1.Push(2)
	s1.Push(3)
	s1.Push(4)
	s1.Push(5)
	return s1
}

func TestReverseStack1(t *testing.T) {
	type args struct {
		myStack *data_structure.Stack
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"1",
			args{GetTestStack()},
		},
	}
	for _, tt := range tests {
		ReverseStack1(tt.args.myStack)
		fmt.Println("逆序栈:")
		for !tt.args.myStack.IsEmpty(){
			tmp ,_:= tt.args.myStack.Pop()
			fmt.Printf("%d ",tmp )
		}
	}
}

func TestReverseStack(t *testing.T) {
	type args struct {
		myStack *data_structure.Stack
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"1",
			args{GetTestStack()},
		},
	}
	for _, tt := range tests {
		ReverseStack(tt.args.myStack)
		fmt.Println("逆序栈:")
		for !tt.args.myStack.IsEmpty(){
			tmp ,_:= tt.args.myStack.Pop()
			fmt.Printf("%d ",tmp )
		}
	}
}
