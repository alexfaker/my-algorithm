package dynamic_planning

func QueNum1(n int) int {
	if n < 1 {
		return 0
	}
	var record = make([]int, n, n)
	return qNumProcess1(0, record, n)
}

func qNumProcess1(i int, record []int, n int) int {
	if i == n {
		// 找到了一种解法
		return 1
	}
	var res = 0

	for j := 0; j < n; j++ {
		if isValid(record, i, j) {
			record[i] = j
			res += qNumProcess1(i+1, record, n)
		}
	}
	return res

}

func isValid(record []int, i, j int) bool {
	for k := 0; k < i; k++ {
		if j == record[k] || abs(record[k]-j) == abs(i-k) {
			return false
		}
	}
	return true
}

func abs(i int) int {
	if i > 0 {
		return i
	}
	return -i
}
