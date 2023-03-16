package tree

import (
	"fmt"
	"strconv"
	"strings"

	"gitee.com/Anderson/my-algorithm/data_structure"
)

// 序列化和反序列化二叉树

// 前序遍历序列化 和反序列化

// 前序序列化二叉树

func PreSerial(head *data_structure.BinaryTreeNode) string {
	if head == nil {
		return "#!"
	}
	res := fmt.Sprintf("%d!", head.Value)
	res += PreSerial(head.Left)
	res += PreSerial(head.Right)
	return res
}

// 反序列化 二叉树

func RevertPreSerial(res string) *data_structure.BinaryTreeNode {
	ans := strings.Split(res, "!")
	if len(ans) == 0 {
		return nil
	}
	var dp = data_structure.NewDoubleDirectQueueStrImp()
	for i := 0; i < len(ans); i++ {
		if len(ans[i]) > 0 {
			dp.RightIn(ans[i])
		}
	}

	return revertPreSerial(dp)
}

func revertPreSerial(ans *data_structure.DoubleDirectQueueStrImp) *data_structure.BinaryTreeNode {
	tmp := ans.LeftOut()
	if tmp == "#" || len(tmp) == 0 {
		return nil
	}
	res, _ := strconv.Atoi(tmp)
	newNode := data_structure.NewBinaryTreeNode(res)
	newNode.Left = revertPreSerial(ans)
	newNode.Right = revertPreSerial(ans)
	return newNode
}

// 层遍历 序列化二叉树

func LevelSerial(head *data_structure.BinaryTreeNode) string {
	if head == nil {
		return "#!"
	}
	var res string
	var dp = []*data_structure.BinaryTreeNode{head}
	for len(dp) > 0 {
		tmp := dp[0]
		if tmp != nil {
			res += fmt.Sprintf("%d!", tmp.Value)
			dp = append(dp, tmp.Left)
			dp = append(dp, tmp.Right)
		} else {
			res += "#!"
		}
		if len(dp) > 1 {
			dp = dp[1:]
		} else {
			dp = []*data_structure.BinaryTreeNode{}
		}
	}

	return res
}

func RevertLevelSerial(res string) *data_structure.BinaryTreeNode {
	if len(res) == 0 {
		return nil
	}
	ans := strings.Split(res, "!")
	if len(ans) == 0 {
		return nil
	}
	var nodeList = []*data_structure.BinaryTreeNode{}

	for i := 0; i < len(ans); i++ {
		if len(ans[i]) > 0 && ans[i] != "#" {
			t, _ := strconv.Atoi(ans[i])
			node := data_structure.NewBinaryTreeNode(t)
			nodeList = append(nodeList, node)
		}
	}
	var head = nodeList[0]
	var cur = 0
	var curNode = head
	for i := 1; i < len(nodeList); i++ {
		if nodeList[i] != nil && curNode.Left == nil {
			curNode.Left = nodeList[i]
			continue
		}
		if nodeList[i] != nil && curNode.Right == nil {
			curNode.Right = nodeList[i]
			cur++
			curNode = nodeList[cur]
		}
	}
	return head
}
