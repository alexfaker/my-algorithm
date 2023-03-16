package dynamic_planning

/*
经典的斐波那契数列问题
1. 最简单的解法为递归
f(1) = 1; f(2) = 1; f(3) = 2
f(n) = f(n-1) + f(n-2)
时间复杂度 O(2^N), 很拉胯

2. 稍微好点的解法为从左到右循环计算每一项的值,计算到第 N 项为止
时间复杂度 O(N)

3. 牛逼的实现, 时间复杂度可以到 O(log(N))
直接说关键结论.
{F(n),F(n-1))} = {F(n-1),F(n-2)} * |1,1|^(n-1)
								   |1,0|
推理: 递归式严格遵循 F(n) = F(n-1) + F(n-2), 那么, 一定有一个矩阵(2*2), 使得
{F(n),F(n-1))} = {F(n-1),F(n-2)} * |x|
带入F(1) = 1 ; F(2) = 1; F(3) = 2; F(4) = 3 解得
|x| = |1, 1|
 	  |1, 0|
那么. 这个问题就转化成了一个如何快速求得(2*2)矩阵的(n) 次方的问题

*/

func Fibonacci1(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	return Fibonacci1(n-1) + Fibonacci1(n-2)
}

func Fibonacci2(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	var pre = 1
	var res = 1
	var tmp = 0
	for i := 3; i <= n; i++ {
		tmp = res
		res += pre
		pre = tmp
	}
	return res
}

func Fibonacci3(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	base := [][]int{
		{1, 1},
		{1, 0},
	}
	res := matrixPower(base, n-2)
	return res[0][0] + res[1][0]
}

func matrixPower(m [][]int, p int) [][]int {
	var res = make([][]int, len(m), len(m))
	for i := 0; i < len(m); i++ {
		res[i] = make([]int, len(m[0]), len(m[0]))
	}
	// 把 res 设为单位矩阵,相当于整数中的 i
	for i := 0; i < len(res); i++ {
		res[i][i] = 1
	}
	tmp := m
	for ; p != 0; p >>= 1 {
		if (p & 1) != 0 {
			res = muliMatrix(res, tmp)
		}
		tmp = muliMatrix(tmp, tmp)
	}
	return res
}

/*
矩阵相乘
公式记忆; 行 列 式  (这里仅为记忆方便, 不要于行列式定义混淆)
比如 矩阵 (M*N) * (N*P) 所得结果矩阵为 M*P
其中 M 的第 a 行 与 N 的 第 B 列 逐项相乘 在相加, 就是 结果矩阵的第 a 行 第 b 列的值
*/
func muliMatrix(m1 [][]int, m2 [][]int) [][]int {
	var res = make([][]int, len(m1), len(m1))
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, len(m2))
	}

	for i := 0; i < len(m2[0]); i++ {
		for j := 0; j < len(m1); j++ {
			for k := 0; k < len(m2); k++ {
				res[i][j] += m1[i][k] * m2[k][j]
			}
		}
	}
	return res
}


func muliMatrix1(m1 [][]int, m2 [][]int) [][]int {
	var res = make([][]int, len(m1), len(m1))
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, len(m2))
	}

	// res[i][j] = m1[i][0-n] * m2[0-n][i]
	for i := 0  ;i < len(m1);i++{
		for j:= 0 ; j < len(m2);j++{
			for k := 0 ; k < len(m1);k++{
				res[i][j] += m1[i][k] * m2[k][j]
			}
		}
	}
	return res
}

/*
补充练习题
农场中成熟的母牛每年只会生一头小母牛,并且永远不会死,
第一年农场有一只成熟的母牛, 从第二年开始, 母牛开始生小母牛.
每只小母牛三年之后成熟有可以开始生小母牛,
给定整数 N,求出 N 年后 牛的数量

F(N) = F(N-1) + F(N-3) n1 =1 n2 = 2
三阶递推公式, 待补充学习相关数学知识
 */

