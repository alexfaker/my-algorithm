package tree

import (
	"gitee.com/Anderson/my-algorithm/data_structure"
)

//已知二叉树的每个结点的值都不一样
//根据二叉树的先序遍历,中序遍历, 后序遍历中的两者遍历数组, 重建这颗二叉树, 并返回这颗树的头节点

/*
通过二叉树的先序遍历数组,中序遍历数组重建二叉树, 并返回二叉树的头节点
思路
1. 定义: 先序遍历数组为 pre , 中序遍历数组为 in
2. 当前先序遍历数组里的第一个节点一定是这棵树的头节点. 没问题吧
3. 在中序遍历数组中, 头节点所在位置的左边, 一定是该头节点的左子树节点集合,  右边, 一定是该头节点的右子树节点集合, 没问题吧
4. 动态规划, 分解重建的规模, 即可解决这个问题

递归函数 重建二叉树,返回头节点 入参(头节点先序遍历数组范围, 头节点中序遍历数组范围)
如何确定先序遍历数组,中序遍历数组范围,用原始数组的索引标记即可, 所以我们还需要一个记录一个所有中序遍历数组索引的数据结构, 比如 map
最终的递归函数为 重建二叉树(先序遍历数组, 先序起始索引, 先序结束索引, 中续遍历数组, 中序起始索引, 中序结束索引, 中序 值 -> 数组索引 map )
*/

func RebuildByPreAndIn(pre, in []int) *data_structure.BinaryTreeNode {
	if len(pre) <= 0 || len(in) <= 0 {
		return nil
	}
	var indexMap = make(map[int]int)
	for i := 0; i < len(in); i++ {
		indexMap[in[i]] = i
	}
	return preIn(pre, 0, len(pre)-1, in, 0, len(in)-1, indexMap)
}

func preIn(pre []int, pi, pj int, in []int, ni, nj int, indexMap map[int]int) *data_structure.BinaryTreeNode {
	if pi > pj {
		return nil
	}
	head := data_structure.NewBinaryTreeNode(pre[pi])
	index := indexMap[pre[pi]]
	head.Left = preIn(pre, pi+1, pi+index-ni, in, ni, index-1, indexMap)
	head.Right = preIn(pre, pi+index-ni+1, pj, in, index+1, nj, indexMap)
	return head
}

/*
中序和后续重建二叉树
*/

func RebuildByInAndBehind(in, after []int) *data_structure.BinaryTreeNode {
	if len(after) <= 0 || len(in) <= 0 {
		return nil
	}
	var indexMap = make(map[int]int)
	for i := 0; i < len(in); i++ {
		indexMap[in[i]] = i
	}
	return inBehind(after, 0, len(after)-1, in, 0, len(in)-1, indexMap)
}

func inBehind(after []int, ai, aj int, in []int, ni, nj int, indexMap map[int]int) *data_structure.BinaryTreeNode {
	if ai > aj {
		return nil
	}
	head := data_structure.NewBinaryTreeNode(after[aj])
	index := indexMap[after[aj]]
	head.Left = inBehind(after, ai, aj-nj+index-1, in, ni, index-1, indexMap)
	head.Right = inBehind(after, aj-nj+index, aj-1, in, index+1, nj, indexMap)
	return head
}

/*
中序和后续重建二叉树
注意. 只有一种树可以被中序和后序 数组重建出来:  除叶结点外, 其他的节点均有左孩子和右孩子
简单描述, 每个节点的 孩子数量为 0 或者 2
*/

func RebuildByPreAndBehind(pre, after []int) *data_structure.BinaryTreeNode {
	if len(after) <= 0 || len(pre) <= 0 {
		return nil
	}
	var indexMap = make(map[int]int)
	for i := 0; i < len(after); i++ {
		indexMap[after[i]] = i
	}
	return preBehind(pre, 0, len(pre)-1, after, 0, len(after)-1, indexMap)
}

func preBehind(pre []int, pi, pj int, after []int, ai, aj int, indexMap map[int]int) *data_structure.BinaryTreeNode {
	head := data_structure.NewBinaryTreeNode(after[aj])
	aj -= 1
	//todo 这个判断是干啥的??? 已经是叶节点了,不需要再重建子树了???
	if pi == pj {
		return head
	}
	pi += 1
	index := indexMap[pre[pi]] //前序遍历 左子树的第一个节点的位置, 在后序遍历中的位置中, 恰巧是左子树的最后一个位置
	head.Left = preBehind(pre, pi, pi+index-ai, after, ai, index, indexMap)
	head.Right = preBehind(pre, pi+index-ai+1, pj, after, index+1, aj, indexMap)
	return head
}
