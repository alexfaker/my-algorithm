package tree

import (
	"fmt"

	"gitee.com/Anderson/my-algorithm/data_structure"
)

/*
二叉树的遍历问题
1. 分别用递归、非递归的方法实现二叉树的前序遍历、中序遍历、后续遍历
*/

/*
递归的方式实现
*/

func PreOrder(head *data_structure.BinaryTreeNode) {
	if head == nil {
		return
	}
	fmt.Printf("%v ", head.Value)
	PreOrder(head.Left)
	PreOrder(head.Right)
}

func MidOrder(head *data_structure.BinaryTreeNode) {
	if head == nil {
		return
	}
	MidOrder(head.Left)
	fmt.Printf("%v ", head.Value)
	MidOrder(head.Right)
}

func AfterOrder(head *data_structure.BinaryTreeNode) {
	if head == nil {
		return
	}
	AfterOrder(head.Left)
	AfterOrder(head.Right)
	fmt.Printf("%v ", head.Value)
}

/*非递归版本的前序遍历
1. 申请一个新的栈 stack && 将头节点head压入栈中
2. 从stack弹出栈顶元素 cur,将cur的右孩子(如果有的话)压入栈中,将cur的左孩子(如果有的话) 压入栈中
3. 不断重复步骤 2 ,直到 stack 为空
*/

func PreOrder1(head *data_structure.BinaryTreeNode) {
	if head == nil {
		return
	}
	var myStack = data_structure.NewAllStack()
	myStack.Push(head)

	for !myStack.IsEmpty() {
		cur, err := myStack.Pop()
		if err != nil {
			return
		}
		node := cur.(*data_structure.BinaryTreeNode)
		fmt.Printf("%v ", node.Value)

		if node.Right != nil {
			myStack.Push(node.Right)
		}
		if node.Left != nil {
			myStack.Push(node.Left)
		}
	}
}

/*非递归版本的中序遍历
1. 申请一个新的栈 stack && 令 cur = headr
2. 把 cur 压入栈中, 对于 cur 这棵树来说, 不断地把该树的左节点压入栈中, 即: 不断地cur = cur.left
3. 不断地重复步骤 2 ,直到 cur 为空, 此时从栈 stack 中弹出一个元素 node ,打印 node,并让 cur = node.right, 然后重复步骤 2
4. stack 为空且 cur 为空 -> 终止
*/

func MidOrder1(head *data_structure.BinaryTreeNode) {
	if head == nil {
		return
	}
	var myStack = data_structure.NewAllStack()
	var cur = head
	for cur != nil || !myStack.IsEmpty() {
		if cur != nil {
			myStack.Push(cur)
			cur = cur.Left
		} else {
			res, err := myStack.Pop()
			if err != nil {
				return
			}
			node := res.(*data_structure.BinaryTreeNode)
			fmt.Printf("%v ", node.Value)
			cur = node.Right
		}
	}
}

/*非递归版本的后续遍历 -> 使用两个栈的实现
1. 申请两个栈 s1, s2, 将头节点压入 s1 中
2. 从s1 中弹出的节点记为 cur,然后依次将 cur 的左孩子和右孩子压入s1 中
3. 每一次从s1 中弹出的元素放入 s2 中
4. 不断重复步骤 2 和步骤 3 直到 s1 为空为止
5. 从 s2 中依次弹出元素并打印,打印结果就是二叉树的后续遍历
*/

func AfterOrder1(head *data_structure.BinaryTreeNode) {
	if head == nil {
		return
	}
	var s1 = data_structure.NewAllStack()
	var s2 = data_structure.NewAllStack()
	s1.Push(head)
	for !s1.IsEmpty() {
		res, err := s1.Pop()
		if err != nil {
			return
		}
		s2.Push(res)
		cur := res.(*data_structure.BinaryTreeNode)
		if cur.Left != nil {
			s1.Push(cur.Left)
		}
		if cur.Right != nil {
			s1.Push(cur.Right)
		}
	}

	// 打印
	for !s2.IsEmpty() {
		res, err := s2.Pop()
		if err != nil {
			return
		}
		cur := res.(*data_structure.BinaryTreeNode)
		fmt.Printf("%v ", cur.Value)
	}
}

/*
遍历升级, 二叉树的神级遍历
要求时间复杂度 O(N), 额外空间复杂度O(1)!
思路:
不能申请额外的空间, 只能利用二叉树节点现有的指针(叶子结点的空指针)
让叶子结点的空指针指向上一层的节点
*/

/*
中序遍历的实现
1. 头节点为 h, 让 h 左子树的最右节点的 右指针指向 h, 然后 h 的左子树继续处理步骤一, 直到某个节点没有左子树, 标记为 node. 进入步骤二
2. 从 node 开始 通过每个指针的 right 指针移动 并以此打印 , 假设移动到的节点为 cur , 判断 cur 的左子树的最右节点的 right 指针是否指向 cur, 有两种情况
	1. 如果是, 让那个 cur 的左子树的最右节点的 right 指针 指向 nil(还原二叉树), 打印 cur, cur = cur.right
 	2. 如果不是, 以 cur 为头的子树重回步骤一 执行
*/

func MorrisIn(head *data_structure.BinaryTreeNode) {
	if head == nil {
		return
	}

	var cur1 = head
	var cur2 *data_structure.BinaryTreeNode

	for cur1 != nil {
		cur2 = cur1.Left
		if cur2 != nil {
			// 找到cur1左子树的最右节点
			for cur2.Right != nil && cur2.Right != cur1 {
				cur2 = cur2.Right
			}

			// 如果最右节点的 right 指针为空, 则指向 cur1
			if cur2.Right == nil {
				cur2.Right = cur1
				cur1 = cur1.Left
				continue
			} else {
				//还原叶子节点的指针
				cur2.Right = nil
			}
		}
		fmt.Printf("%v ", cur1.Value)
		cur1 = cur1.Right
	}
	println()
}

/*
先序遍历的实现
整体思路和中序遍历一致, 只要调整打印时机就可以了
在第一次找到节点 h 的左子树最右节点 || 左子树最右节点为空时, 打印该节点
*/

func MorrisPre(head *data_structure.BinaryTreeNode) {
	if head == nil {
		return
	}

	var cur1 = head
	var cur2 *data_structure.BinaryTreeNode

	for cur1 != nil {
		cur2 = cur1.Left
		if cur2 != nil {
			// 找到cur1左子树的最右节点
			for cur2.Right != nil && cur2.Right != cur1 {
				cur2 = cur2.Right
			}

			// 如果最右节点的 right 指针为空, 则指向 cur1
			if cur2.Right == nil {
				cur2.Right = cur1
				fmt.Printf("%v ", cur1.Value)
				cur1 = cur1.Left
				continue
			} else {
				//还原叶子节点的指针
				cur2.Right = nil
			}
		} else {
			fmt.Printf("%v ", cur1.Value)
		}
		cur1 = cur1.Right
	}
	println()
}

/*
后序遍历的实现
todo  待实现. 太麻烦了
*/

func MorrisAfter(head *data_structure.BinaryTreeNode) {
	if head == nil {
		return
	}

	var cur1 = head
	var cur2 *data_structure.BinaryTreeNode

	for cur1 != nil {
		cur2 = cur1.Left
		if cur2 != nil {
			// 找到cur1左子树的最右节点
			for cur2.Right != nil && cur2.Right != cur1 {
				cur2 = cur2.Right
			}

			// 如果最右节点的 right 指针为空, 则指向 cur1
			if cur2.Right == nil {
				cur2.Right = cur1
				fmt.Printf("%v ", cur1.Value)
				cur1 = cur1.Left
				continue
			} else {
				//还原叶子节点的指针
				cur2.Right = nil
			}
		} else {
			fmt.Printf("%v ", cur1.Value)
		}
		cur1 = cur1.Right
	}
	println()
}


// 高阶 二叉树的神级遍历方法
// 实现二叉树的前序遍历, 中序遍历, 后续遍历,
// 要求时间复杂度 O(N), 空间复杂度 O(1)

// 先序遍历的实现
// func MorrisPre(head *data_structure.BinaryTreeNode) {

// }
