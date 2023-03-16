package dynamic_planning

func SubSets(arr []int) [][]int {
	n := len(arr)
	var res [][]int
	for i := 0; i < 1<<n; i++ {
		set := []int{}
		for j := 0; j < n; j++ {
			if i>>j&1 > 0 {
				set = append(set, arr[j])
			}
		}
		res = append(res, append([]int(nil), set...))
	}
	return res
}
