package stack

import (
	"fmt"
	"math"

	"gitee.com/Anderson/my-algorithm/data_structure"
)

//汉诺塔问题 不能跨越一个桩移动
//1. 递归实现
//2. 栈实现

//func main() {
//	fmt.Printf("total move cnt:%d\n", hanoProblem2(3, "left", "right", "middle"))
//}

func hanoProblem1() {

}

type MoveAction int

const (
	No            MoveAction = 0
	LeftToMiddle  MoveAction = 1
	MiddleToLeft  MoveAction = 2
	MiddleToRight MoveAction = 3
	RightToMiddle MoveAction = 4
)

func hanoProblem2(num int, left, right, middle string) int {
	LS := data_structure.NewStack()
	MS := data_structure.NewStack()
	RS := data_structure.NewStack()
	LS.Push(math.MaxInt64)
	RS.Push(math.MaxInt64)
	MS.Push(math.MaxInt64)
	step := 0
	for i := num; i > 0; i-- {
		LS.Push(i)
	}
	action := []MoveAction{No}
	for RS.Size() != num+1 {
		//一直循环操作
		step += fStackTotStack(action, MiddleToLeft, LeftToMiddle, LS, MS, left, middle)
		step += fStackTotStack(action, LeftToMiddle, MiddleToLeft, MS, LS, middle, left)
		step += fStackTotStack(action, MiddleToRight, RightToMiddle, RS, MS, right, middle)
		step += fStackTotStack(action, RightToMiddle, MiddleToRight, MS, RS, middle, right)
	}
	return step
}

func fStackTotStack(action []MoveAction, preAct MoveAction, curAct MoveAction,
	fStack *data_structure.Stack, tStack *data_structure.Stack, from, to string) int {
	fTop, _ := fStack.Peek()
	tTop, _ := tStack.Peek()
	if action[0] != preAct && fTop < tTop {
		tmp, _ := fStack.Pop()
		tStack.Push(tmp)
		fmt.Printf("move %d from %s to %s\n", tmp, from, to)
		action[0] = curAct
		return 1
	}
	return 0
}
