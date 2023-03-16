package arr_and_matrix

type NumMatrix struct {
	Matrix [][]int
}


func Constructor(matrix [][]int) NumMatrix {
	var m = len(matrix)
	var n = len(matrix[0])
	var dp = make([][]int,m,m)
	for i := 0; i< m;i++{
		dp[i] = make([]int,n,n)
	}
	dp[0][0] = matrix[0][0]

	for i := 1; i < m;i++{
		dp[i][0] = dp[i-1][0] + matrix[i][0]
	}

	for j := 1; j < n;j++{
		dp[0][j] = dp[0][j-1] + matrix[0][j]
	}
	for i := 1; i < m ; i++{
		for j := 1; j < n ;j++{
			dp[i][j] = dp[i-1][j] + dp[i][j-1] + matrix[i][j] - dp[i-1][j-1]
		}
	}

	return NumMatrix{dp}

}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	var res int
	res = this.Matrix[row2][col2]
	if col1 -1 >= 0 {
		res -= this.Matrix[row2][col1-1]
	}
	if row1-1 >= 0 {
		res -=  this.Matrix[row1-1][col2]
	}
	if col1 -1 >= 0 && row1-1 >= 0 {
		res +=  this.Matrix[row1-1][col1-1]
	}
	return res

}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */