package main

import (
	"fmt"
)

// 复制带随机指针的链表
// func main() {
//	var head = data_structure.NewRandStackNode(1)
//	var n1 = data_structure.NewRandStackNode(3)
//	var n2 = data_structure.NewRandStackNode(2)
//	var n3 = data_structure.NewRandStackNode(1)
//	head.Next = n1
//	head.Rand = n3
//	n1.Next = n2
//	n1.Rand = n1
//	n2.Next = n3
//	n2.Rand = n3
//	n3.Rand = n2
//
//	newHead := stack.CopyRandNode(head)
//	for cur := newHead; cur != nil; cur = cur.Next {
//		fmt.Printf("new nodelist curNode val:%v,rand:%+v\n", cur.Value, cur.Rand)
//	}
//
//	for cur := head; cur != nil; cur = cur.Next {
//		fmt.Printf("old nodelist curNode val:%v,rand:%+v\n", cur.Value, cur.Rand)
//	}
// }

// 打印两个有序链表的公共部分
// func main() {
//	var l1 = node_list.NewNode(1)
//	var l2 = node_list.NewNode(2)
//	var l3 = node_list.NewNode(3)
//	var l4 = node_list.NewNode(5)
//	var l5 = node_list.NewNode(7)
//	l1.Next = l2
//	l2.Next = l3
//	l3.Next = l4
//	l4.Next = l5
//
//	var m1 = node_list.NewNode(1)
//	var m2 = node_list.NewNode(2)
//	var m3 = node_list.NewNode(4)
//	var m4 = node_list.NewNode(5)
//	var m5 = node_list.NewNode(9)
//	var m6 = node_list.NewNode(11)
//	m1.Next = m2
//	m2.Next = m3
//	m3.Next = m4
//	m4.Next = m5
//	m5.Next = m6
//
//	node_list.PrintCommon(l1,m1)
// }

// 逆序单链表
// func main() {
//	var l1 = node_list.NewNode(1)
//	var l2 = node_list.NewNode(2)
//	var l3 = node_list.NewNode(3)
//	var l4 = node_list.NewNode(5)
//	var l5 = node_list.NewNode(7)
//	l1.Next = l2
//	l2.Next = l3
//	l3.Next = l4
//	l4.Next = l5
//	newHead := node_list.ReverseNode(l1)
//	var cur = newHead
//	for cur != nil {
//		fmt.Println(cur.Value)
//		cur = cur.Next
//	}
//
// }

// 约瑟夫问题
// func main() {
//	var l1 = node_list.NewNode(1)
//	var l2 = node_list.NewNode(2)
//	var l3 = node_list.NewNode(3)
//	var l4 = node_list.NewNode(4)
//	var l5 = node_list.NewNode(5)
//	var l6 = node_list.NewNode(6)
//	var l7 = node_list.NewNode(7)
//	l1.Next = l2
//	l2.Next = l3
//	l3.Next = l4
//	l4.Next = l5
//	l5.Next = l6
//	l6.Next = l7
//	l7.Next = l1
//	res := node_list.JosePusKill(l1, 3)
//	fmt.Printf("the one :%v", res.Value)
// }

// 判断回文链表

// func main() {
//	var l1 = node_list.NewNode(1)
//	var l2 = node_list.NewNode(2)
//	var l3 = node_list.NewNode(3)
//	var l4 = node_list.NewNode(2)
//	var l5 = node_list.NewNode(1)
//	l1.Next = l2
//	l2.Next = l3
//	l3.Next = l4
//	l4.Next = l5
//
//	fmt.Println(node_list.IsPalindrome(l1))
//	fmt.Println(node_list.IsPalindrome2(l1))
//
//	var l6 = node_list.NewNode(1)
//	var l7 = node_list.NewNode(2)
//	var l8 = node_list.NewNode(2)
//	var l9 = node_list.NewNode(1)
//	l6.Next = l7
//	l7.Next = l8
//	l8.Next = l9
//	fmt.Println(node_list.IsPalindrome(l6))
//	fmt.Println(node_list.IsPalindrome2(l6))
//
//	var m1 = node_list.NewNode(1)
//	var m2 = node_list.NewNode(2)
//	var m3 = node_list.NewNode(4)
//	var m4 = node_list.NewNode(5)
//	m1.Next = m2
//	m2.Next = m3
//	m3.Next = m4
//	fmt.Println(node_list.IsPalindrome(m1))
//	fmt.Println(node_list.IsPalindrome2(m1))
// }

/*
搜索二叉树 中序遍历 转双向链表
*/

// func main() {
//	var l1 = node_list.NewBinaryTreeNode(6)
//	var l2 = node_list.NewBinaryTreeNode(4)
//	var l3 = node_list.NewBinaryTreeNode(7)
//	var l4 = node_list.NewBinaryTreeNode(2)
//	var l5 = node_list.NewBinaryTreeNode(5)
//	var l6 = node_list.NewBinaryTreeNode(9)
//	var l7 = node_list.NewBinaryTreeNode(1)
//	var l8 = node_list.NewBinaryTreeNode(3)
//	var l9 = node_list.NewBinaryTreeNode(8)
//	l1.Left = l2
//	l1.Right = l3
//	l2.Left = l4
//	l2.Right = l5
//	l3.Right = l6
//	l4.Left = l7
//	l4.Right = l8
//	l6.Left = l9
//
//	res := node_list.Convert1(l1)
//	var tmp = res
//	var last = res
//	for tmp != nil {
//		fmt.Println(tmp.Value)
//		if tmp.Next == nil {
//			last = tmp
//		}
//		tmp = tmp.Next
//	}
//	fmt.Println("逆序")
//
//	for last != nil {
//		fmt.Println(last.Value)
//		last = last.Last
//	}
//
// }

// 获取栈内的最小元素
// func main() {
//	var res =stack.NewMinStackImp()
//	res.Push(21)
//	fmt.Println(res.GetMin())
//	res.Push(12)
//	fmt.Println(res.GetMin())
//	res.Push(3)
//	fmt.Println(res.GetMin())
//	res.Pop()
//	fmt.Println(res.GetMin())
// }

// 栈的逆序

// func main() {
//	myStack := &data_structure.Stack{List: []int{5, 4, 3, 2, 1}}
//
//	stack.ReverseStack(myStack)
//	fmt.Println(myStack.Pop())
//	fmt.Println(myStack.Pop())
//	fmt.Println(myStack.Pop())
//	fmt.Println(myStack.Pop())
//	fmt.Println(myStack.Pop())
//
// }

// 栈实现一个队列

// func main() {
//	var res = stack.NewQueueByStack()
//	res.Add(1)
//	res.Add(2)
//	res.Add(3)
//	fmt.Println(res.Poll())
//	fmt.Println(res.Poll())
//	fmt.Println(res.Poll())
//	res.Add(4)
//	res.Add(5)
//	res.Add(6)
//	fmt.Println(res.Poll())
//	fmt.Println(res.Poll())
//	res.Add(7)
//	fmt.Println(res.Peek())
// }

// 二叉树的遍历问题
// func main() {
//	var l1 = data_structure.NewBinaryTreeNode(6)
//	var l2 = data_structure.NewBinaryTreeNode(4)
//	var l3 = data_structure.NewBinaryTreeNode(7)
//	var l4 = data_structure.NewBinaryTreeNode(2)
//	var l5 = data_structure.NewBinaryTreeNode(5)
//	var l6 = data_structure.NewBinaryTreeNode(9)
//	var l7 = data_structure.NewBinaryTreeNode(1)
//	var l8 = data_structure.NewBinaryTreeNode(3)
//	var l9 = data_structure.NewBinaryTreeNode(8)
//	l1.Left = l2
//	l1.Right = l3
//	l2.Left = l4
//	l2.Right = l5
//	l3.Right = l6
//	l4.Left = l7
//	l4.Right = l8
//	l6.Left = l9
//
//	fmt.Printf("递  归:前序 ")
//	tree.PreOrder(l1)
//	fmt.Println()
//
//	fmt.Printf("非递归:前序 ")
//	tree.PreOrder1(l1)
//	fmt.Println()
//
//	fmt.Printf("MorrisIn:前序 ")
//	tree.MorrisPre(l1)
//	fmt.Println()
//	//
//	//fmt.Printf("递  归:中序 ")
//	//tree.MidOrder(l1)
//	//fmt.Println()
//	//
//	//fmt.Printf("非递归:中序 ")
//	//tree.MidOrder1(l1)
//	//fmt.Println()
//	//fmt.Printf("MorrisIn:中序 ")
//	//tree.MorrisIn(l1)
//	//fmt.Println()
//
//
//
//	//fmt.Printf("递  归:后序 ")
//	//tree.AfterOrder(l1)
//	//fmt.Println()
//	//
//	//fmt.Printf("非递归:后序 ")
//	//tree.AfterOrder(l1)
//
// }

/*
求无序数组的最长和为 k 的子数组
*/

// func main() {
//	//var arr = []int{1, 2, 1, 1}
//	//fmt.Println(arr_and_matrix.GetMaxLength(arr, 3))
//	//fmt.Println(arr_and_matrix.GetMaxLength(arr, 2))
//	//fmt.Println(arr_and_matrix.GetMaxLength(arr, 10))
//
//	var arr1 = []int{1, 2, 1, 1, 6}
//	fmt.Println(arr_and_matrix.GetMaxLength(arr1, 10))
//	fmt.Println(arr_and_matrix.GetMaxLength1(arr1, 10))
//	fmt.Println(arr_and_matrix.GetMaxLength2(arr1, 10))
// }

// 二叉树中两个节点的最小公共祖先

// func main() {
//	var l1 = data_structure.NewBinaryTreeNode(6)
//	var l2 = data_structure.NewBinaryTreeNode(4)
//	var l3 = data_structure.NewBinaryTreeNode(7)
//	var l4 = data_structure.NewBinaryTreeNode(2)
//	var l5 = data_structure.NewBinaryTreeNode(5)
//	var l6 = data_structure.NewBinaryTreeNode(9)
//	var l7 = data_structure.NewBinaryTreeNode(1)
//	var l8 = data_structure.NewBinaryTreeNode(3)
//	var l9 = data_structure.NewBinaryTreeNode(8)
//	l1.Left = l2
//	l1.Right = l3
//	l2.Left = l4
//	l2.Right = l5
//	l3.Right = l6
//	l4.Left = l7
//	l4.Right = l8
//	l6.Left = l9
//
//	o1 := l3
//	o2 := l4
//	fmt.Printf("动态规划版本 res:%+v\n", tree.LowestAncestor(l1, o1, o2))
//	fmt.Printf("额外空间版本 res:%+v\n", tree.LowestAncestor1(l1, o1, o2))
//
// }

/*
中序遍历数组和先序遍历数组重建二叉树
*/

// func main() {
//	pre := []int{1, 2, 4, 5, 8, 9, 3, 6, 7}
//	in := []int{4, 2, 8, 5, 9, 1, 6, 3, 7}
//	after := []int{4, 8, 9, 5, 2, 6, 7, 3, 1}
//
//	head := tree.RebuildByPreAndIn(pre, in)
//	tree.PreOrder1(head)
//	println()
//	tree.MidOrder1(head)
//	println()
//
//	head1 := tree.RebuildByInAndBehind(in, after)
//	tree.MidOrder1(head1)
//	println()
//	tree.AfterOrder1(head1)
//	println()
//
//	head2 := tree.RebuildByPreAndBehind(pre, after)
//	tree.PreOrder1(head2)
//	println()
//	tree.AfterOrder1(head2)
//
// }

/*
统计和生成所有不同的二叉树
*/
// func main() {
// 	fmt.Println(tree.GetAllCntByMiddleOrder(3))
// }

/*
二叉树节点间的最大距离
*/

// func main() {
// 	var l1 = data_structure.NewBinaryTreeNode(6)
// 	var l2 = data_structure.NewBinaryTreeNode(4)
// 	var l3 = data_structure.NewBinaryTreeNode(7)
// 	var l4 = data_structure.NewBinaryTreeNode(2)
// 	var l5 = data_structure.NewBinaryTreeNode(5)
// 	var l6 = data_structure.NewBinaryTreeNode(9)
// 	var l7 = data_structure.NewBinaryTreeNode(1)
// 	var l8 = data_structure.NewBinaryTreeNode(3)
// 	var l9 = data_structure.NewBinaryTreeNode(8)
// 	l1.Left = l2
// 	l1.Right = l3
// 	l2.Left = l4
// 	l2.Right = l5
// 	l3.Right = l6
// 	l4.Left = l7
// 	l4.Right = l8
// 	l6.Left = l9
// 	fmt.Println(tree.MaxDistance(l1))
//
// 	var l1a = data_structure.NewBinaryTreeNode(1)
// 	var l2a = data_structure.NewBinaryTreeNode(2)
// 	var l3a = data_structure.NewBinaryTreeNode(3)
// 	var l4a = data_structure.NewBinaryTreeNode(4)
// 	var l5a = data_structure.NewBinaryTreeNode(5)
// 	var l6a = data_structure.NewBinaryTreeNode(6)
// 	var l7a = data_structure.NewBinaryTreeNode(7)
// 	l1a.Left = l2a
// 	l1a.Right = l3a
// 	l2a.Left = l4a
// 	l2a.Right = l5a
// 	l3a.Left = l6a
// 	l3a.Right = l7a
//
// 	fmt.Println(tree.MaxDistance(l1a))
// }

/*
统计完全二叉树的节点个数
要求 时间复杂度小于O(N)
 */

// func main()  {
// 		var l1a = data_structure.NewBinaryTreeNode(1)
// 		var l2a = data_structure.NewBinaryTreeNode(2)
// 		var l3a = data_structure.NewBinaryTreeNode(3)
// 		var l4a = data_structure.NewBinaryTreeNode(4)
// 		var l5a = data_structure.NewBinaryTreeNode(5)
// 		var l6a = data_structure.NewBinaryTreeNode(6)
// 		var l7a = data_structure.NewBinaryTreeNode(7)
// 		l1a.Left = l2a
// 		l1a.Right = l3a
// 		l2a.Left = l4a
// 		l2a.Right = l5a
// 		l3a.Left = l6a
// 		l3a.Right = l7a
//
// 		println(tree.CBTNodeNum(l1a))
// }

/*
斐波那契数列 O(logN) 时间复杂度的实现
*/
// func main()  {
// 	fmt.Println(dynamic_planning.Fibonacci3(1))
// 	fmt.Println(dynamic_planning.Fibonacci3(2))
// 	fmt.Println(dynamic_planning.Fibonacci3(3))
// 	fmt.Println(dynamic_planning.Fibonacci3(89))
// }

/*
最少硬币数问题
*/
// func main() {
//  // coins := []int{7, 2, 5}
//  // fmt.Println(dynamic_planning.MinCoins(coins, 27))
//  // fmt.Println(dynamic_planning.MinCoins1(coins, 27))
//
//  // 扩展问题, 最多硬币数问题
//  coins1 := []int{5, 10, 25, 1}
//  fmt.Println(dynamic_planning.MaxCoins1(coins1, 15))
//  return
// }

/*
机器人走网格所有走法问题
*/

// func main() {
//  var res = make([][]int, 2, 2)
//  for i := 0; i < len(res); i++ {
//   res[i] = make([]int, len(res), len(res))
//  }
//  fmt.Println(dynamic_planning.MinMoveSteps(res))
//  return
// }

/*
 矩阵的最小路径和
 */
// func main() {
// 	var arr = [][]int{
// 		{1, 3, 5, 9},
// 		{8, 1, 3, 4},
// 		{5, 0, 6, 1},
// 		{8, 8, 4, 0},
// 	}
// 	fmt.Println(dynamic_planning.GetMinPathSum(arr))
// 	fmt.Println(dynamic_planning.GetMinPathSum1(arr))
//
// 	var arr1 = [][]int{
// 		{1, 3, 5, 9},
// 		{8, 1, 3, 4},
// 		{5, 0, 6, 1},
// 		{8, 8, 4, 0},
// 		{1, 2, 4, 7},
// 	}
// 	fmt.Println(dynamic_planning.GetMinPathSum(arr1))
// 	fmt.Println(dynamic_planning.GetMinPathSum1(arr1))
//
// 	var arr2 = [][]int{
// 		{1, 3, 5, 9},
// 		{8, 1, 3, 4},
// 		{5, 0, 6, 1},
// 	}
// 	fmt.Println(dynamic_planning.GetMinPathSum(arr2))
// 	fmt.Println(dynamic_planning.GetMinPathSum1(arr2))
// }

// func main(){
// 	var arr = []int{2,1,5,3,4,8,9,7}
// 	fmt.Println(dynamic_planning.GetLongestSubArr(arr))
// }

// func main(){
// 	fmt.Println(dynamic_planning.QueNum1(8))
// }

// 排序算法
// func main(){
// 	// var arr = []int{7,8,3,5,9}
// 	// fmt.Println(my_sort.BubbleSort(arr))
// 	// fmt.Println(my_sort.QuickSort(arr))
//
//
// 	var arr1 = []int{7,8,3,5,9,1,12,34,23,656,10}
// 	// fmt.Println(my_sort.BubbleSort(arr1))
// 	// fmt.Println(my_sort.QuickSort(arr1))
// 	// fmt.Println(my_sort.InsertSort(arr1))
// 	// fmt.Println(my_sort.InsertSort1(arr1))
// 	// fmt.Println(my_sort.MergeSort(arr1))
// 	fmt.Println(my_sort.MergeSort1(arr1))
// }

/*
数组的最大子数组和 递归实现 && 动态规划的实现
*/
// func main() {
// 	// var arr = []int{7, 1, -3, 5, -9}
// 	// fmt.Println(dynamic_planning.MaxSubArr(arr))
// 	// fmt.Println(dynamic_planning.MaxSubArr1(arr))
// 	// var arr1 = []int{7, 1, 3, 5, -9}
// 	// fmt.Println(dynamic_planning.MaxSubArr(arr1))
// 	// fmt.Println(dynamic_planning.MaxSubArr1(arr1))
//
// 	var arr2 = []int{7, 1, 3, 5, 9}
//
// 	fmt.Println(my_sort.CountingSort(arr2, 10))
// }

// func main(){
// 	var prices = []int{1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
// 	// fmt.Println(dynamic_planning.MaxCutRodProfit(prices, 4))
// 	// fmt.Println(dynamic_planning.MaxCutRodProfit1(prices, 4))
// 	// fmt.Println(dynamic_planning.MaxCutRodProfit2(prices, 4))
// 	// fmt.Println(dynamic_planning.MaxCutRodProfit(prices, 10))
// 	// fmt.Println(dynamic_planning.MaxCutRodProfit1(prices, 10))
// 	// fmt.Println(dynamic_planning.MaxCutRodProfit2(prices, 10))
// 	maxPro, res := dynamic_planning.MaxCutRodProfit3(prices, 7)
// 	fmt.Println(maxPro)
// 	n := 7
// 	for n > 0 {
// 		fmt.Println(res[n])
// 		n -= res[n]
// 	}
// }

// func main() {
// 	var str1 = "abcde"
// 	var str2 = "abdecc"
// 	fmt.Println(strings.IsDeformation(str1, str2))
//
// 	var str3 = "abcde"
// 	var str4 = "abdec"
// 	fmt.Println(strings.IsDeformation(str3, str4))
// }
//
// func main() {
// 	var str1 = "ab21cd--32-e"
// 	fmt.Println(strings.NumOfChar(str1))
//
// }

func main() {
	// var matrix = [][]int{
	// 	{-4,-5},
	// }
	// var res = arr_and_matrix.Constructor(matrix)
	// fmt.Println(res.SumRegion(0,0,0,1))
	// fmt.Println(countSegments(", , , ,        a, eaefa"))
	fmt.Println(reverseStr("abcdefg", 2))

}

func reverseStr(s string, k int) string {
	var l = len(s)
	var newS = s
	if l == 0 || l < 2*k {
		return newS
	}
	// 反转 前 k 个字符
	newS = revertStrPro(newS, 0, k-1)

	if len(s[2*k:]) <= k {
		newS = revertStrPro(newS, 2*k, l-1)
	} else if len(newS[2*k:]) >= k && len(newS[2*k:]) < 2*k {
		newS = revertStrPro(newS, 2*k, 2*k+k-1)
	}
	return newS

}

func revertStrPro(str string, s, e int) string {
	newStr := []rune(str)
	if s >= e {
		return string(newStr)
	}
	for i := s; i < s+(e-s)/2; i++ {
		newStr[s], newStr[e] = newStr[e], newStr[s]
	}
	return string(newStr)
}