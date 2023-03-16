package dynamic_planning

// 给定两个 字符串 str1 , str2
// 求两个字符串的最长公共子串
// 比如 str1 = A1234, str2 = CD1234
// 最长公共子串为 1234


// 动态规划解法
// 1. 先求出最长公共子串的长度和结尾的字符
// dp[i][j] = 以str1[i] 和 str1[j] 为结尾的最长公共子串长度
// 如果 str1[i] != str2[j], 那么说明以str1[i] 和 str1[j]结尾没有公共子串, 长度为 0
// 如果 str1[i] == str2[j], 那么说明以str1[i] 和 str1[j]结尾有公共子串, 长度为dp[i-1][j-1]+1, 也就是左上角的长度+1
// 2. 根据最长子串的长度, 和结尾字符对应的索引, 反推出最长公共子串
// 时间复杂度 O(M*N), m,n 分别为 str1 , str2 的长度
// 空间复杂度 O(M*N)
func getLongestSubStr(subStr1, subStr2 string) string {
	var res string
	var l1, l2 = len(subStr1), len(subStr2)
	if l1 == 0 || l2 == 0 {
		return ""
	}

	// 初始化 dp
	var dp = make([][]int, l1, l1)
	for i := 0; i < l1; i++ {
		dp[i] = make([]int, l2, l2)
	}
	// 初始化第一行和第一列
	for i := 0; i < l2; i++ {
		if subStr1[0] == subStr2[i] {
			dp[0][i] = 1
		}
	}
	for i := 0; i < l1; i++ {
		if subStr1[i] == subStr2[0] {
			dp[i][0] = 1
		}
	}
	for i := 1; i < l1; i++ {
		for j := 1; j < l2; j++ {
			if subStr1[i] == subStr2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
		}
	}

	// 找到最大值

	var maxlen = 0
	var l1Index = 0
	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			if dp[i][j] > maxlen {
				maxlen = dp[i][j]
				l1Index = i
			}
		}
	}
	if maxlen == 0 {
		return ""
	}
	res = subStr1[maxlen-l1Index+1 : l1Index+1]
	return res
}

// 优化版本, 空间复杂度O(1)
// 如解法一所描述, 其实只要知道 dp 矩阵的左上角的值, 就可以求得当前值了
// 斜着求就可以求出来 最终的最长子串长度 和索引位置 牛逼 plus
func getLongestSubStrV2(subStr1, subStr2 string) string {
	var res string
	var l1, l2 = len(subStr1), len(subStr2)
	if l1 == 0 || l2 == 0 {
		return ""
	}
	var row = 0
	var col = l2 - 1
	var maxLen = 0
	var index = 0
	for row < l1 {
		i := row
		j := col
		ml := 0
		for i < l1 && j < l2 {
			if subStr1[i] == subStr2[j] {
				ml++
			}
			if ml > maxLen {
				maxLen = ml
				index = i
			}
			i++
			j++
		}
		if col > 0 {
			col--
		} else {
			row++
		}
	}

	// 找到最大值
	if maxLen == 0 {
		return ""
	}

	res = subStr1[maxLen-index+1 : index+1]
	return res
}
