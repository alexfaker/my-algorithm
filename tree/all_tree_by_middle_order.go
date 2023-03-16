package tree

/*
给定一个整数 N, 如果 N < 1 ,代表空树结构, 否则代表 中序遍历的结果为 {1,2,3,4,5}
请返回可能得二叉树结构有多少个
 */
//  待理解这个 O(N) 的实现?
func GetAllCntByMiddleOrder(n int) int {
	if n <= 1 {
		return n
	}
	var num = make([]int, n+1, n+1)
	num[0] = 1
	// todo  这里为什么这样处理
	for i := 1; i <= n; i++ {
		for j := 1; j < i+1; j++ {
			num[i] += num[j-1] * num[i-j]
		}
	}
	// 统计可能的树结构
	return num[n]
}
