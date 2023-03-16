package dynamic_planning

import (
	"fmt"
	"math"
)

/*
给定一个数组, 代表硬币的面值(不重复), 再给定一个值 aim 代表钱数
求凑出指定钱数的最小硬币个数
比如 硬币面值 res = [2,5,7], aim = 27

凑出 27 需要的最小硬币数为 7* 1 + 5 * 4  = 5 个
*/

func MinCoins(coins []int, aim int) int {
	if len(coins) <= 0 || aim <= 0 {
		return -1
	}
	var n = len(coins)
	var max = math.MaxInt32
	var dp = make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, aim+1)
	}

	for j := 1; j <= aim; j++ {
		dp[0][j] = max
		if j-coins[0] >= 0 && dp[0][j-coins[0]] != max {
			dp[0][j] = dp[0][j-coins[0]] + 1
		}
	}
	var left int
	for i := 1; i < n; i++ {
		for j := 1; j <= aim; j++ {
			left = max
			if j-coins[i] >= 0 && dp[i][j-coins[i]] != max {
				left = dp[i][j-coins[i]] + 1
			}
			dp[i][j] = int(math.Min(float64(left), float64(dp[i-1][j])))
		}
	}
	if dp[n-1][aim] != max {
		return dp[n-1][aim]
	}
	return -1
}

/*
焯 ,这个实现简单多了
动态规划四部曲
1. 确定状态(1. 最后一步 2. 转化成子问题)
2. 转移方程
3. 开始和边界条件
4. 计算顺序

以这一道题为例
给定条件: 货币面值为[2,5,7], 要找的钱总共是 27 元
第一步:确定状态(最后一步 && 子问题)
	* 假设最优解需要 k 张货币, 那个前面的货币面值是 27-k
	* 那么, 考虑最后一步, 一定是这最后一张货币 加入之后, 凑成了 27 元, 且达到了最优解
	* 上一步的钱数为 27 -k, 这个问题的规模比原规模小了(确定子问题),接下来我们考虑怎么求解子问题
    * 假设函数 f(x) = 用最少得硬币拼出 x 的数量. x 为钱总数
	* 那么对于最后一步来说, 最后那枚硬币只可能是 2 5 7 中的一个
		* 如果最后一张是 2 , f(27) = f(27-2) + 1 (前面的所有硬币数 + 最后一枚 2 )
		* 如果最后一张是 5 , f(27) = f(27-5) + 1 (前面的所有硬币数 + 最后一枚 5 )
		* 如果最后一张是 7 , f(27) = f(27-7) + 1 (前面的所有硬币数 + 最后一枚 7 )
	* 这样子问题就确定了
第二步: 确定转移方程
	* 有了上一步的铺垫, 我们可以轻松的写出 f(x) 这个方程, 我们用一个数组res[x] 来记录每一个子问题的记录
	* x 表示要兑换的总面值. val 表示最小硬币数
	* res[x] = min{res[x-2] + 1, res[x-5]+1,res[x-7]+1}
第三部: 确定边界条件
	* 如果 x-2 || x-5 || x-7 小于 0 怎么办?
		* ???
	* 所有的硬币 拼不出来给的总面值怎么办?
		* 拼不出来, 就将数组对应元素设置为正无穷, 或者最大正数
	* 初始条件怎么定义?
		* res[0] = 0 , 表示面值为 0 时, 需要 0个硬币
*/

// func MinCoins1(coins []int, aim int) int {
// 	if len(coins) <= 0 || aim <= 0 {
// 		return -1
// 	}
// 	var res = make([]int, aim+1, aim+1)
// 	res[0] = 0 // 这一步在 go 里是多余的, 但是还是写出来,便于理解
// 	for i := 1; i <= aim; i++ {
// 		res[i] = math.MaxInt32 // 默认值, 表示不能凑出该面值
// 		for j := 0; j < len(coins); j++ {
// 			// 这个逻辑能写出来就神了, 妙哉妙哉!
// 			// 这里感觉带不带 res[i-coins[j]] != math.MaxInt64都可以,
// 			// 但是为了防止比较是数据类型转换的异常, 还是带上比较好
// 			if i-coins[j] >= 0 && res[i-coins[j]] != math.MaxInt64 {
// 				res[i] = int(math.Min(float64(res[i]), float64(res[i-coins[j]]+1)))
// 			}
// 		}
// 	}
// 	if res[aim] == math.MaxInt64 {
// 		return -1
// 	}
// 	return res[aim]
// }

func MinCoins1(coins []int, aim int) int {
	if len(coins) <= 0 || aim <= 0 {
		return -1
	}
	var res = make([]int, aim+1, aim+1)
	res[0] = 0 // 这一步在 go 里是多余的, 但是还是写出来,便于理解
	for i := 1; i <= aim; i++ {
		res[i] = math.MaxInt32 // 默认值, 表示不能凑出该面值
		for j := 0; j < len(coins); j++ {
			// 这个逻辑能写出来就神了, 妙哉妙哉!
			// 这里感觉带不带 res[i-coins[j]] != math.MaxInt64都可以,
			// 但是为了防止比较是数据类型转换的异常, 还是带上比较好
			if i-coins[j] >= 0 && res[i-coins[j]] != math.MaxInt64 {
				if res[i] > res[i-coins[j]]+1 {
					res[i] = res[i-coins[j]] + 1
					if i == aim {
						fmt.Printf("%d ", coins[j])
					}
				} else {
					res[i] = res[i]
				}
			}
		}
	}
	if res[aim] == math.MaxInt64 {
		return -1
	}
	return res[aim]
}


/*
 扩展问题, 求换钱总共有多少种方法
这里实现4 种解法
1. 暴力递归法  时间复杂度 O(aim^N), 空间复杂度O(1)
2. 基于暴力递归的记忆优化 时间复杂度 O(N), 空间复杂度O(1)
3. 动态规划 时间复杂度 O(N), 空间复杂度O(1)
4. 枚举过程优化后的动态规划 时间复杂度 O(N), 空间复杂度O(1)
5. 枚举过程优化后&& 压缩使用空间的动态规划 时间复杂度 O(N), 空间复杂度O(1)
*/

/*
解法1 : 暴力递归实现
如果 coins = {5,10,25,1}, aim = 1000 ,分析过程如下
1. 用 0 个 5 元的硬币, 让[10,25,1] 组成剩下的 1000, 所有方法数记为 res1
2. 用 1 个 5 元的硬币, 让[10,25,1] 组成剩下的 995, 所有方法数记为 res2
3. 用 2 个 5 元的硬币, 让[10,25,1] 组成剩下的 990, 所有方法数记为 res3
......
201. 用 200 个 5 元的硬币, 让[10,25,1] 组成剩下的 0, 所有方法数记为 res201
res1+res2+...res201 就是总的方法数
*/

func MaxCoins1(coins []int, aim int) int {
	if len(coins) == 0 || aim <= 0 {
		return 0
	}
	return process1(coins, 0, aim)
}

func process1(arr []int, index, aim int) int {
	var res = 0
	if index == len(arr) {
		if aim == 0 {
			res = 1
		}
	} else {
		for i := 0; arr[index]*i <= aim; i++ {
			res += process1(arr, index+1, aim-arr[index]*i)
		}
	}
	return res
}


func MaxCoins3(coins []int, aim int) int {
	if len(coins) == 0 || aim <= 0 {
		return 0
	}
	var dp = make([][]int, len(coins), len(coins))
	for i := 0; i < len(coins); i++ {
		dp[i] = make([]int, aim+1, aim+1)
	}

	// 边界条件初始化, 每种面值 换 0 元, 都只有一种方式
	for i := 0; i < len(coins); i++ {
		dp[i][0] = 1
	}
	// 只有第一种面值, 可以换的钱
	for j := 1; coins[0]*j <= aim; j++ {
		dp[0][coins[0]*j] = 1
	}

	for i := 1; i < len(coins); i++ {
		for j := 1; j <= aim; j++ {
			num := 0
			for k := 0; j-coins[i]*k >= 0; k++ {
				num += dp[i-1][j-coins[i]*k]
			}
			dp[i][j] = num
		}
	}
	return dp[len(coins)-1][aim]
}

// 时间复杂度 O(N*aim)的实现

func MaxCoins4(coins []int, aim int) int {
	// 边界条件处理
	var n = len(coins)
	if n == 0 || aim == 0 {
		return -1
	}
	var dp = make([][]int, n, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, aim+1, aim+1)
	}
	// 初始化动态规划边界条件
	for i := 0; i < n; i++ {
		dp[i][0] = 1
	}

	for i := 0; i <= aim; i++ {
		for k := 1; k*coins[0] <= aim; k++ {
			dp[0][k*coins[0]] = 1
		}
	}
	// for j := 1; coins[0]*j <= aim; j++ {
	// 	dp[0][coins[0]*j] = 1
	// }
	// 动态规划数组顺序计算
	for i := 1; i < n; i++ {
		for j := 1; j <= aim; j++ {
			dp[i][j] = dp[i-1][j]
			if j-coins[i] >= 0 {
				dp[i][j] += dp[i][j-coins[i]]
			}
		}
	}
	return dp[n-1][aim]
}

